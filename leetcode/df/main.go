package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Amenity struct {
	Name     string `json:"name"`
	SeoName  string `json:"seo_name"`
	ImageURL string `json:"image_url"`
}

type Outlet struct {
	OutletID           int64     `json:"outlet_id"`
	SiteID             int64     `json:"site_id"`
	OutletName         string    `json:"outlet_name"`
	OpeningHours       string    `json:"opening_hours"`
	Directions         string    `json:"directions"`
	City               string    `json:"city"`
	SiteName           string    `json:"site_name"`
	Country            string    `json:"country"`
	ServiceDescription string    `json:"service_description"`
	TerminalName       string    `json:"terminal_name"`
	TerminalType       string    `json:"terminal_type"`
	GateType           string    `json:"gate_type"`
	PricingGroup       *string   `json:"pricing_group"`
	ServiceID          int       `json:"service_id"`
	Amenities          []Amenity `json:"amenities"`
}

type Response struct {
	Status     bool     `json:"status"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       []Outlet `json:"data"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("dfresponse.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON data
	var response Response
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Count the number of outlets
	outletCount := len(response.Data)
	fmt.Printf("Number of outlets: %d\n", outletCount)
}
