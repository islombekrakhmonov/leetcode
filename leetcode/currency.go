package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
)

type Currency struct {
	Title string  `json:"title" example:"US Dollar"`
	Price float64 `json:"cb_price,string" example:"0.000000000000000000"`
}

func main() {
	// url := "https://nbu.uz/en/exchange-rates/json/"

	// response, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println("Error fetching data:", err)
	// 	return
	// }
	// defer response.Body.Close()

	// var exchangeRates []Currency
	// err = json.NewDecoder(response.Body).Decode(&exchangeRates)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }

	// var usd Currency
	// for _, rate := range exchangeRates {
	// 	if rate.Title == "US Dollar" {
	// 		usd = rate
	// 	}
	// }

	// fmt.Println("US Dollar:", usd)

	// currentTime := time.Now()
	// expirationTime := calculateExpirationTime(currentTime)
	inputTimes := []time.Time{
		// Monday at 10 am
		time.Date(2023, time.November, 13, 10, 0, 0, 0, time.UTC),
		// Wednesday at 3 pm
		time.Date(2023, time.November, 15, 15, 0, 0, 0, time.UTC),
		// Friday at 5 pm
		time.Date(2023, time.November, 17, 17, 0, 0, 0, time.UTC),
		// Saturday at 11 am
		time.Date(2023, time.November, 18, 11, 0, 0, 0, time.UTC),
		// Sunday at 2 pm
		time.Date(2023, time.November, 19, 14, 0, 0, 0, time.UTC),
	}

	for _, inputTime := range inputTimes {
		expirationTime := calculateExpirationTime(inputTime)

		fmt.Printf("Input Time: %s\n", inputTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("Expiration Time: %s\n", expirationTime.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}
}
func calculateExpirationTime(startTime time.Time) time.Time {
	businessCal := cal.NewBusinessCalendar()

	businessCal.SetWorkHours(time.Duration(8)*time.Hour, time.Duration(16)*time.Hour)

	if startTime.Weekday() == time.Saturday {
		startTime = businessCal.NextWorkdayStart(startTime)
	} else if startTime.Weekday() == time.Sunday {
		startTime = businessCal.NextWorkdayStart(startTime)
	}

	if startTime.Hour() >= 16 {
		startTime = businessCal.NextWorkdayStart(startTime)
	}

	expirationTime := businessCal.WorkdaysFrom(startTime, 4)

	expirationTime = businessCal.WorkdayEnd(expirationTime)

	return expirationTime
}
