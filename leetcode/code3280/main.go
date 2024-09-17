package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(convertDateToBinary("2080-02-29"))
}

func convertDateToBinary(date string) string {

	var output string
	dateSlice := strings.Split(date, "-")

	for _, value := range dateSlice {
		number, _ := strconv.Atoi(value)
		binaryRepresentation := strconv.FormatInt(int64(number), 2)

		output += binaryRepresentation
	}

	return output[:len(output)-1]
}
