package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(winningPlayer(2, 7))
	// fmt.Println(Add10Minutes("2025-07-02", "-10:00"))
	// fmt.Println(IsCurrentDate("2025-07-03", "03:00"))
	fmt.Println(CalculateAvailableAt(25, "03:00", false))
}

// min 1 x and 4 y
func winningPlayer(x int, y int) string {
	turn := 0

	for {
		if (x <= 0 && y <= 3) || x < 1 || y < 4 {
			break
		}

		turn++
		x--
		y -= 4
	}

	if turn%2 == 0 {
		return "Bob"
	}

	return "Alice"
}

func Add10MinutesIfToday(serviceDate string, timezoneOffset string) (string, error) {

	date, err := time.Parse(time.DateOnly, serviceDate)
	if err != nil {
		return "", err
	}

	timeZoneOffsetStringHour := strings.Split(timezoneOffset, ":")[0] + "h"
	timeZoneOffsetStringMinute := strings.Split(timezoneOffset, ":")[1] + "m"

	timeZoneOffset, err := time.ParseDuration(timeZoneOffsetStringHour + timeZoneOffsetStringMinute)
	if err != nil {
		return "", err
	}

	return date.Add(timeZoneOffset + 10*time.Minute).Format(time.DateTime), nil
}

func IsCurrentDate(serviceDate, timezoneOffset string) (bool, error) {

	date, err := time.Parse(time.DateOnly, serviceDate)
	if err != nil {
		return false, err
	}

	timeZoneOffsetStringHour := strings.Split(timezoneOffset, ":")[0] + "h"
	timeZoneOffsetStringMinute := strings.Split(timezoneOffset, ":")[1] + "m"

	timeZoneOffset, err := time.ParseDuration(timeZoneOffsetStringHour + timeZoneOffsetStringMinute)
	if err != nil {
		return false, err
	}

	return time.Now().Add(timeZoneOffset).Format(time.DateOnly) == date.Format(time.DateOnly), nil
}

func CalculateAvailableAt(preOrderTime int, timezoneOffset string, IsFlighRequired bool) (availableAt string) {
	if preOrderTime < 24 && preOrderTime > 0 {
		IsFlighRequired = true
	}

	// Parse timezone offset
	parts := strings.Split(timezoneOffset, ":")
	if len(parts) != 2 {
		return ""
	}

	hoursPart := parts[0] + "h"
	minutesPart := parts[1] + "m"

	offsetHours, err1 := time.ParseDuration(hoursPart)
	offsetMinutes, err2 := time.ParseDuration(minutesPart)
	if err1 != nil || err2 != nil {
		return ""
	}

	timeZoneOffset := offsetHours + offsetMinutes
	now := time.Now().UTC().Add(timeZoneOffset)

	if IsFlighRequired || preOrderTime == 0 {
		availableAt = now.Add(time.Hour * time.Duration(preOrderTime)).Format(time.DateTime)
	} else {
		calculatedTime := time.Now().Add(timeZoneOffset).Add(time.Hour * time.Duration(preOrderTime))
		calculatedTime = calculatedTime.Add(24 * time.Hour)
		midnight := time.Date(calculatedTime.Year(), calculatedTime.Month(), calculatedTime.Day(), 0, 0, 0, 0, time.UTC)

		availableAt = midnight.Format(time.DateTime)
	}

	return availableAt
}
