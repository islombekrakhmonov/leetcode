package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
	//Output: "2052-10-20"

}

func reformatDate(date string) string {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	monthToNum := make(map[string]string)

	for i, v := range months {
		var iStr string
		if i < 9 {
			iStr += "0"
		}
		iStr += strconv.Itoa(i + 1)
		monthToNum[v] = iStr
	}

	splitted := strings.Split(date, " ")

	var output string

	output += splitted[2] + "-"

	month := splitted[1]
	dateOut := splitted[0]

	monthToOutput := monthToNum[month]
	dateOut = dateOut[:len(dateOut)-2]
	dateOutInt, _ := strconv.Atoi(dateOut)
	if dateOutInt < 10 {
		dateOut = "0" + dateOut
	}
	output += monthToOutput + "-" + dateOut

	return output
}
