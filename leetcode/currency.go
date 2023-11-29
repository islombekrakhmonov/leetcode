package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gammban/numtow"
	"github.com/gammban/numtow/lang"
	"github.com/rickar/cal/v2"
)

var (
	units     = []string{"ноль", "один", "два", "три", "четыре", "пять", "шесть", "семь", "воcемь", "девять"}
	teens     = []string{"", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens      = []string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds  = []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
	thousands = []string{"", "тысяча", "тысячи", "тысяч"}
	millions  = []string{"", "миллион", "миллиона", "миллионов"}
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

	// usd := 0.0

	// for _, rate := range exchangeRates {
	// 	if rate.Title == "US Dollar" {
	// 		usd = rate.Price
	// 	}
	// }

	// fmt.Println("US Dollar:", usd)

	// // currentTime := time.Now()
	// // expirationTime := calculateExpirationTime(currentTime)
	// inputTimes := []time.Time{
	// 	// Monday at 10 am
	// 	time.Date(2023, time.November, 13, 10, 0, 0, 0, time.UTC),
	// 	// Wednesday at 3 pm
	// 	time.Date(2023, time.November, 15, 15, 0, 0, 0, time.UTC),
	// 	// Friday at 5 pm
	// 	time.Date(2023, time.November, 17, 17, 0, 0, 0, time.UTC),
	// 	// Saturday at 11 am
	// 	time.Date(2023, time.November, 18, 11, 0, 0, 0, time.UTC),
	// 	// Sunday at 2 pm
	// 	time.Date(2023, time.November, 19, 14, 0, 0, 0, time.UTC),
	// }

	// for _, inputTime := range inputTimes {
	// 	expirationTime := calculateExpirationTime(inputTime)

	// 	fmt.Printf("Input Time: %s\n", inputTime.Format("2006-01-02 15:04:05"))
	// 	fmt.Printf("Expiration Time: %s\n", expirationTime.Format("2006-01-02 15:04:05"))
	// 	fmt.Println()
	// }

	price := 470.67
	usd := 12304.98

	multiply := price * usd
	myString := fmt.Sprintf("%.2f", multiply)
	fmt.Println(myString)
	intPart := int(multiply)

	currency := "uzs" 
	convertedString := numtow.MustString(myString, lang.RU)
	splitted := strings.Split(convertedString, " ")

	for i,v := range splitted {
		if v == "сотых" {
			if currency == "usd" {
				splitted[i] = "цент"
			} else if currency == "uzs" {
				splitted[i] = "тийинов"
			}
		} else if v == "целых" {
			if currency == "usd" {
				if intPart%10 > 1 && intPart%10 < 5 {
					splitted[i] = "доллара"
				} else {
					splitted[i] = "доллар"
				} 	
			} else if currency == "uzs" {
				splitted[i] = "сум"
			}
		}
	}
}


func ConvertNumberToRussianText(amount float64, currency string) string {
	var (
		currencyUnit string
		decimalUnit  string
	)

	fmt.Printf("%.2f\n", amount)
	intPart := int(amount)
	fmt.Println(intPart)
	decimalPart := Round(float64(amount-float64(intPart)), 2) // round to 2 decimal places
	decimalInt := int(decimalPart * 100)
	fmt.Println(decimalInt)

	intPartText := toRussianWords(intPart)
	fmt.Println(intPartText)
	decimalPartText := toRussianWords(decimalInt)

	if currency == "usd" {
		if intPart%10 > 1 && intPart%10 < 5 {
			currencyUnit = "доллара"
		} else {
			currencyUnit = "доллар"
		}

		if decimalInt%10 > 1 && decimalInt%10 < 5 {
			decimalUnit = "цента"
		} else {
			decimalUnit = "цент"
		}
	} else if currency == "uzs" {
		if intPart%10 > 1 && intPart%10 < 5 {
			currencyUnit = "сума"
		} else {
			currencyUnit = "сум"
		}

		if decimalInt%10 > 1 && decimalInt%10 < 5 {
			decimalUnit = "тийина"
		} else {
			decimalUnit = "тийин"
		}
	}

	var result string
	if intPartText == "" {
		result = decimalPartText + " " + decimalUnit
	} else if decimalPartText == "" {
		result = intPartText + " " + currencyUnit
	} else {
		result = fmt.Sprintf("%s %s %s %s", intPartText, currencyUnit, decimalPartText, decimalUnit)
	}

	return strings.TrimSpace(result)
}

func Round(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(x*pow) / pow
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

func toRussianWords(number int) string {
	if number < 0 {
		return ""
	} else if number == 0 {
		return "ноль"
	}

	var result string
	var scales = []struct {
		exp  int
		name string
	}{
		{6, "миллион"},
		{3, "тысяча"},
	}

	for _, scale := range scales {
		scaleValue := int(math.Pow10(scale.exp))
		if number >= scaleValue {
			count := number / scaleValue
			number %= scaleValue
			if result != "" {
				result += " "
			}
			result += toRussianWords(count) + " " + scale.name
			if count > 1 && count < 5 {
				result += "и"
			} else if count >= 5 || count == 0 {
				result += ""
			}
		}
	}

	for number > 0 {
		if result != "" {
			result += " "
		}
		if number < 10 {
			result += units[number]
			break
		} else if number < 20 {
			result += teens[number-10]
			break
		} else if number < 100 {
			result += fmt.Sprintf("%s %s", tens[number/10], units[number%10])
			break
		} else {
			thousands := number / 1000
			if thousands > 0 {
				result += fmt.Sprintf("%s тысяч", toRussianWords(thousands))
				if thousands > 1 && thousands < 5 {
					result += "и"
				} else if thousands >= 5 || thousands == 0 {
					result += ""
				}
			}
			number %= 1000
		}
	}

	return result
}