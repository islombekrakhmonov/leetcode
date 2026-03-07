package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasAlternatingBits(n int) bool {

	binStr := strconv.FormatInt(int64(n), 2)

	for i := 0; i < len(binStr)-1; i++ {
		if binStr[i] == binStr[i+1] {
			return false
		}
	}

	return true
}

type CancellationFee struct {
	FeeInPercents              int `json:"feeInPercents"`
	HoursBeforeServiceDateTime int `json:"hoursBeforeServiceDateTime"`
}

func ParseCancellationFees(input string) ([]CancellationFee, error) {
	rules := strings.Split(input, ";")
	var fees []CancellationFee

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		parts := strings.Split(rule, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid rule format: %s", rule)
		}

		// left side: hours (e.g. "48h")
		hoursStr := strings.TrimSpace(parts[0])
		hoursStr = strings.TrimSuffix(hoursStr, "h")
		hours, err := strconv.Atoi(hoursStr)
		if err != nil {
			return nil, fmt.Errorf("invalid hours in rule %s: %w", rule, err)
		}

		// right side: percent (e.g. "50%")
		percentStr := strings.TrimSpace(parts[1])
		percentStr = strings.TrimSuffix(percentStr, "%")
		percent, err := strconv.Atoi(percentStr)
		if err != nil {
			return nil, fmt.Errorf("invalid percent in rule %s: %w", rule, err)
		}

		fees = append(fees, CancellationFee{
			FeeInPercents:              percent,
			HoursBeforeServiceDateTime: hours,
		})
	}

	return fees, nil
}

func main() {
	input := "48h - 50%; 24h - 70%; 2h - 100%"
	fees, err := ParseCancellationFees(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fees)
}
