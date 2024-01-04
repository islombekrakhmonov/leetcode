package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
}

func areNumbersAscending(s string) bool {
	splitted := strings.Split(s, " ")
	lastNum := -1
    
    for i := 0; i < len(splitted); i++ {
        if value, err := strconv.Atoi(splitted[i]); err == nil {
            if lastNum != -1 && value <= lastNum {
                return false
            }
            
            lastNum = value
        }      
    }
	
	return true
}

func removeNonNumbers(input string) string {
	reg := regexp.MustCompile("[^0-9]")

	result := reg.ReplaceAllString(input, "")

	return result
}
