package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(convertTime("02:30", "04:35"))
}

func convertTime(current string, correct string) int {
	var operations int

	currentMinutes, _ := ConvertTimeToMinutes(current)
	correctMinutes, _ := ConvertTimeToMinutes(correct)

	difference := correctMinutes - currentMinutes

	for difference > 0 {
		if difference-60 >= 0 {
			difference -= 60
		} else if difference-15 >= 0 {
			difference -= 15
		} else if difference-5 >= 0 {
			difference -= 5
		} else {
			difference -= 1
		}
		operations++
	}

	return operations
}

func ConvertTimeToMinutes(timeStr string) (int, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid time format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hour value: %w", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minute value: %w", err)
	}

	totalMinutes := hours*60 + minutes
	return totalMinutes, nil
}
