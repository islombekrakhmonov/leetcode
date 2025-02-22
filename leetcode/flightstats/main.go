package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// GetFlightInfo fetches flight information from the FlightStats API
func GetFlightInfo(baseURL, appID, appKey, departureAirport, arrivalAirport string, departureDate time.Time) ([]FlightStatsResponse, error) {

	queryParams := url.Values{}
	queryParams.Set("appId", appID)
	queryParams.Set("appKey", appKey)

	url := fmt.Sprintf("%s/schedules/rest/v1/json/from/%s/to/%s/departing/%d/%d/%d?%s",
		baseURL, departureAirport, arrivalAirport, departureDate.Year(), departureDate.Month(), departureDate.Day(), queryParams.Encode())

	// Prepare the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response
	var schedules FlightResponse
	err = json.Unmarshal(body, &schedules)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON response: %w", err)
	}

	var response []FlightStatsResponse

	for _, flight := range schedules.ScheduledFlights {
		fmt.Println("Flight Number:", flight.FlightEquipmentIataCode, flight.FlightNumber)
		fmt.Println("Departure Time:", flight.DepartureTime)
		departureTime, err := time.Parse("2006-01-02T15:04:05.000", flight.DepartureTime)
		if err != nil {
			return nil, fmt.Errorf("error parsing departure time: %w", err)
		}

		departureTime = departureTime.UTC()

		formattedDepartureTime := departureTime.Format("2006-01-02T15:04:05Z")
		response = append(response, FlightStatsResponse{
			FlightNumber:  flight.CarrierFsCode + flight.FlightNumber,
			DepartureTime: formattedDepartureTime,
		})
	}

	return response, nil
}

func main() {
	// Replace with your actual credentials and desired parameters
	baseURL := "https://api.flightstats.com/flex"
	appID := "49ebeb7d"
	appKey := "32e95eefc8be1be11b564465fc46081d"
	departureAirport := "ICN"
	arrivalAirport := "TAS"

	departureDate := "2024-12-19"

	// Convert the date string to a time.Time object
	departureDateConverted, err := time.Parse("2006-01-02", departureDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// Get flight information
	response, err := GetFlightInfo(baseURL, appID, appKey, departureAirport, arrivalAirport, departureDateConverted)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheduled Flights:", len(response))

}

type FlightStatsResponse struct {
	FlightNumber  string `json:"flightNumber"`
	DepartureTime string `json:"departureTime"`
}

type FlightResponse struct {
	Request          Request           `json:"request"`
	ScheduledFlights []ScheduledFlight `json:"scheduledFlights"`
	Appendix         Appendix          `json:"appendix"`
}

type Request struct {
	Departing        bool   `json:"departing"`
	URL              string `json:"url"`
	DepartureAirport struct {
		RequestedCode string `json:"requestedCode"`
		FsCode        string `json:"fsCode"`
	} `json:"departureAirport"`
	ArrivalAirport struct {
		RequestedCode string `json:"requestedCode"`
		FsCode        string `json:"fsCode"`
	} `json:"arrivalAirport"`
	Date struct {
		Year        string `json:"year"`
		Month       string `json:"month"`
		Day         string `json:"day"`
		Interpreted string `json:"interpreted"`
	} `json:"date"`
}

type Appendix struct {
	Airlines   []Airline   `json:"airlines"`
	Airports   []Airport   `json:"airports"`
	Equipments []Equipment `json:"equipments"`
}

type ScheduledFlight struct {
	CarrierFsCode           string        `json:"carrierFsCode"`
	FlightNumber            string        `json:"flightNumber"`
	DepartureAirportFsCode  string        `json:"departureAirportFsCode"`
	ArrivalAirportFsCode    string        `json:"arrivalAirportFsCode"`
	DepartureTime           string        `json:"departureTime"`
	ArrivalTime             string        `json:"arrivalTime"`
	Stops                   int           `json:"stops"`
	DepartureTerminal       string        `json:"departureTerminal"`
	ArrivalTerminal         string        `json:"arrivalTerminal"`
	FlightEquipmentIataCode string        `json:"flightEquipmentIataCode"`
	IsCodeshare             bool          `json:"isCodeshare"`
	IsWetlease              bool          `json:"isWetlease"`
	ServiceType             string        `json:"serviceType"`
	ServiceClasses          []string      `json:"serviceClasses"`
	TrafficRestrictions     []interface{} `json:"trafficRestrictions"`
	Codeshares              []Codeshare   `json:"codeshares"`
	ReferenceCode           string        `json:"referenceCode"`
	Operator                *Operator     `json:"operator,omitempty"`
}

type Codeshare struct {
	CarrierFsCode       string        `json:"carrierFsCode"`
	FlightNumber        string        `json:"flightNumber"`
	ServiceType         string        `json:"serviceType"`
	ServiceClasses      []string      `json:"serviceClasses"`
	TrafficRestrictions []interface{} `json:"trafficRestrictions"`
	ReferenceCode       int           `json:"referenceCode"`
}

type Operator struct {
	CarrierFsCode       string        `json:"carrierFsCode"`
	FlightNumber        string        `json:"flightNumber"`
	ServiceType         string        `json:"serviceType"`
	ServiceClasses      []string      `json:"serviceClasses"`
	TrafficRestrictions []interface{} `json:"trafficRestrictions"`
}

type Airline struct {
	Fs     string `json:"fs"`
	Iata   string `json:"iata"`
	Icao   string `json:"icao"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Airport struct {
	Fs                 string  `json:"fs"`
	Iata               string  `json:"iata"`
	Icao               string  `json:"icao"`
	Faa                string  `json:"faa"`
	Name               string  `json:"name"`
	City               string  `json:"city"`
	CityCode           string  `json:"cityCode"`
	CountryCode        string  `json:"countryCode"`
	CountryName        string  `json:"countryName"`
	RegionName         string  `json:"regionName"`
	TimeZoneRegionName string  `json:"timeZoneRegionName"`
	WeatherZone        string  `json:"weatherZone"`
	LocalTime          string  `json:"localTime"`
	UtcOffsetHours     float64 `json:"utcOffsetHours"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	ElevationFeet      int     `json:"elevationFeet"`
	Classification     int     `json:"classification"`
	Active             bool    `json:"active"`
}

type Equipment struct {
	Iata      string `json:"iata"`
	Name      string `json:"name"`
	TurboProp bool   `json:"turboProp"`
	Jet       bool   `json:"jet"`
	Widebody  bool   `json:"widebody"`
	Regional  bool   `json:"regional"`
}
