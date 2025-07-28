package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// // fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(ParseFlightNumber(" EK002"))
	fmt.Println(ParseFlightNumber(" EK 02"))
	fmt.Println(ParseFlightNumber("EK 002"))
	fmt.Println(ParseFlightNumber(" EK0022"))
	fmt.Println(ParseFlightNumber("EK02"))
}

func compress(chars []byte) int {

	return 0
}

func ParseFlightNumber(flight string) (string, string, error) {
	flight = strings.ReplaceAll(flight, " ", "")
	flight = strings.TrimSpace(flight)

	// Updated regex: Carrier is 2-3 alphanumeric characters (starting with a letter), optionally ending with '*',
	// followed by 1-4 digits, optionally ending with a letter
	re := regexp.MustCompile(`^([A-Z0-9][A-Z0-9]?\*?)(\d{1,4}[A-Z]?)$`)

	matches := re.FindStringSubmatch(flight)
	if len(matches) < 3 {
		return "", "", fmt.Errorf("invalid flight number format")
	}

	carrier := matches[1]
	number := matches[2]

	number = strings.TrimLeft(number, "0")

	return carrier, number, nil
}
