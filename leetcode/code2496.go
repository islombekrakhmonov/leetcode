package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(maximumValue([]string{"alic3","bob","3","4","00000"}))
}

func maximumValue(strs []string) int {
	var max int 

	for i := 0; i<len(strs); i++{
		if isNumeric(strs[i]) {
			num, _ := strconv.Atoi(strs[i])
			if num > max {
				max = num
			}
		} else {
			if max < len(strs[i]) {
				max = len(strs[i])
			}
		}
	}

	return max
}

func isNumeric(s string) bool {
	var numberRegexString = "^[0-9]+$"

	reg := regexp.MustCompile(numberRegexString)
	
	return reg.MatchString(s)
}
