package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// fmt.Println(bitwiseComplement(5))

func main() {
	tests := []string{"M2*454", "1P*454", "M3454", "MF454", "6E4875", "TEZ1401"}

	for _, flightNumber := range tests {
		carrier, flightNumber, err := ParseFlightNumber(flightNumber)
		if err != nil {
			fmt.Printf("Error parsing flight number %q: %v\n", flightNumber, err)
		} else {
			fmt.Printf("Carrier: %s, Flight Number: %s\n", carrier, flightNumber)
		}
	}
}

func ParseFlightNumber(flight string) (string, string, error) {
	// Updated regex: Carrier is 1-2 alphanumeric characters (starting with a letter or digit), optionally ending with '*',
	// followed by 1-4 digits, optionally ending with a letter
	re := regexp.MustCompile(`^([A-Z0-9][A-Z0-9]?\*?)(\d{1,4}[A-Z]?)$`)

	matches := re.FindStringSubmatch(flight)
	if len(matches) < 3 {
		return "", "", fmt.Errorf("invalid flight number format", flight)
	}

	return matches[1], matches[2], nil
}

func bitwiseComplement(n int) int {

	binary := strconv.FormatInt(int64(n), 2)

	runes := []rune(binary)

	for i, v := range runes {
		if v == '1' {
			runes[i] = '0'
		} else {
			runes[i] = '1'
		}
	}

	binary = string(runes)

	number, _ := strconv.ParseInt(binary, 2, 64)

	return int(number)
}

//n ^ (1<<uint(len(fmt.Sprintf("%b", n))) - 1)
