package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	sdk "github.com/asadbekGo/ett-code-sdk"
	"github.com/spf13/cast"
)

func FlightsBlackList(request sdk.NewRequestBody, ettUcodeApi *sdk.ObjectFunction) ([]byte, sdk.ResponseError) {
	var (
		response      sdk.Response
		errorResponse sdk.ResponseError
		// objectData    = cast.ToStringMap(request.Data["object_data"])
		createMonitoringObjects []map[string]interface{}
	)

	if request.Data["app_id"] == nil {
		errorResponse.StatusCode = 400
		errorResponse.Description = "app_id is required"
		errorResponse.ClientErrorMessage = "app_id is required"
		errorResponse.ErrorMessage = ettUcodeApi.Logger.ErrorLog.Sprint(sdk.ErrorCodeWithMessage[errorResponse.StatusCode])
		return nil, errorResponse
	}
	ettUcodeApi.Cfg.SetAppId(cast.ToString(request.Data["app_id"]))

	getRestrinctions, _, err := ettUcodeApi.GetListSlim(
		&sdk.Argument{
			TableSlug: "restriction_value",
			Request: sdk.Request{
				Data: map[string]interface{}{
					"restriction_type_id": "3115da70-6405-402c-bb8c-05f5267aae9a",
					"limit":               1000000,
				},
			},
			BlockCached: true,
		},
	)

	if err != nil {
		errorResponse.StatusCode = 500
		errorResponse.Description = "Failed to get flight black list data"
		errorResponse.ClientErrorMessage = "Failed to get flight black list data"
		errorResponse.ErrorMessage = ettUcodeApi.Logger.ErrorLog.Sprint(err.Error())
		return nil, errorResponse
	}

	flightNumber := "M2*454"

	carrier, flightNumber, err := ParseFlightNumber(flightNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carrier, flightNumber)

	flightNumber = "1P*454"
	carrier, flightNumber, err = ParseFlightNumber(flightNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carrier, flightNumber)

	flightNumber = "M3454"

	carrier, flightNumber, err = ParseFlightNumber(flightNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carrier, flightNumber)

	flightNumber = "MF454"

	carrier, flightNumber, err = ParseFlightNumber(flightNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(carrier, flightNumber)

	fmt.Println(len(getRestrinctions.Data.Data.Response))

	// airlines, errorResponse := Airlines()
	// if errorResponse.ErrorMessage != "" {
	// 	fmt.Println(errorResponse)
	// }

	file, err := os.Open("airlines.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parse the JSON into the struct
	var airlines AirlinesResponse
	err = json.Unmarshal(fileContent, &airlines)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Println("len(airlines.Airlines):", len(airlines.Airlines))

	fsCodes := make(map[string]string)

	for _, airline := range airlines.Airlines {
		if airline.Iata == "" {
			continue
		}
		_, exists := fsCodes[airline.Iata]
		if exists {
			fmt.Println("Duplicate IATA code:", airline)
		} else {
			fsCodes[airline.Iata] = airline.Fs
		}

	}

	fmt.Println("len(fsCodes):", len(fsCodes))

	getAirlines, _, err := ettUcodeApi.GetListSlim(
		&sdk.Argument{
			TableSlug: "airline",
			Request: sdk.Request{
				Data: map[string]interface{}{
					"limit": 1000000,
				},
			},
			BlockCached: true,
		},
	)

	fmt.Println("len(getAirlines.Data.Data.Response):", len(getAirlines.Data.Data.Response))

	for _, airline := range getAirlines.Data.Data.Response {
		airlineCode := cast.ToString(airline["code"])
		fscode := fsCodes[airlineCode]
		if fscode == "" {
			fmt.Println("No FS code for airline:", airline)
			continue
		}
		guid := cast.ToString(airline["guid"])

		var createMonitoringObject = map[string]interface{}{
			"guid":    guid,
			"fs_code": fscode,
		}

		createMonitoringObjects = append(createMonitoringObjects, createMonitoringObject)
	}

	// fmt.Println(createMonitoringObjects)

	// make json

	// jsonData, err := json.Marshal(airlines)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(jsonData))

	// for _, restriction := range getRestrinctions.Data.Data.Response {
	// 	name := cast.ToString(restriction["description"])
	// 	airlineCode := cast.ToString(restriction["title"])

	// 	name = strings.TrimLeft(name, " ")

	// 	var (
	// 		randUUID, _        = uuid.NewRandom()
	// 		monitoringFlightId = randUUID.String()
	// 	)

	// 	var createMonitoringObject = map[string]interface{}{
	// 		"guid":   monitoringFlightId,
	// 		"code":   airlineCode,
	// 		"name":   name,
	// 		"is_new": true,
	// 	}

	// 	createMonitoringObjects = append(createMonitoringObjects, createMonitoringObject)

	// }

	// if len(createMonitoringObjects) > 0 {
	// 	_, response, err = ettUcodeApi.MultipleUpdate(&sdk.Argument{Ctx: context.Background(), TableSlug: "airline", Request: sdk.Request{Data: map[string]interface{}{"objects": createMonitoringObjects}}, BlockBuilder: true, DisableFaas: true})
	// 	if err != nil {
	// 		errorResponse.StatusCode = 500
	// 		errorResponse.Description = response.Data["description"]
	// 		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[errorResponse.StatusCode]
	// 		errorResponse.ErrorMessage = ettUcodeApi.Logger.ErrorLog.Sprint(err.Error())
	// 		return nil, errorResponse
	// 	}
	// }

	response.Data = map[string]interface{}{}
	response.Status = "done"
	responseByte, _ := json.Marshal(response)

	return responseByte, errorResponse
}
func ParseFlightNumber(flight string) (string, string, error) {
	// Updated regex: Carrier is 2-3 alphanumeric characters (starting with a letter), optionally ending with '*',
	// followed by 1-4 digits, optionally ending with a letter
	re := regexp.MustCompile(`^([A-Z][A-Z0-9]?\*?)(\d{1,4}[A-Z]?)$`)

	matches := re.FindStringSubmatch(flight)
	if len(matches) < 3 {
		return "", "", fmt.Errorf("invalid flight number format")
	}

	return matches[1], matches[2], nil
}

func Airlines() (airlines AirlinesResponse, errorResponse sdk.ResponseError) {

	var (
		queryParams = url.Values{"appId": []string{"611983b1"}, "appKey": []string{"805b6b7a62c9f87942fbf471d750f361"}}
		url         = fmt.Sprintf("%sairlines/rest/v1/json/active?%s", "https://api.flightstats.com/flex/", queryParams.Encode())
	)
	// Prepare the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errorResponse.StatusCode = 500
		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[errorResponse.StatusCode]
		errorResponse.ErrorMessage = err.Error()
		return airlines, errorResponse
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorResponse.StatusCode = 500
		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[errorResponse.StatusCode]
		errorResponse.ErrorMessage = err.Error()
		return airlines, errorResponse
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorResponse.StatusCode = 500
		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[errorResponse.StatusCode]
		errorResponse.ErrorMessage = err.Error()
		return airlines, errorResponse
	}

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		errorResponse.StatusCode = resp.StatusCode
		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[resp.StatusCode]
		errorResponse.ErrorMessage = string(body)
		return airlines, errorResponse
	}

	// Unmarshal the JSON response
	err = json.Unmarshal(body, &airlines)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		errorResponse.StatusCode = 500
		errorResponse.ClientErrorMessage = sdk.ErrorCodeWithMessage[errorResponse.StatusCode]
		errorResponse.ErrorMessage = err.Error()
		return airlines, errorResponse
	}

	return airlines, errorResponse
}

type AirlinesResponse struct {
	Airlines []Airline `json:"airlines"`
}

type Airline struct {
	Fs     string `json:"fs"`
	Iata   string `json:"iata,omitempty"`
	Icao   string `json:"icao"`
	Name   string `json:"name"`
	Active bool   `json:"active"` // Corrected to bool
}
