package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digitCount("1210"))
}

func digitCount(num string) bool {
	var count int 
	for i:=0; i<len(num); i++{
		count = 0
		for j:=0; j<len(num); j++{
			numStr,_ := strconv.Atoi(string(num[j])) 
			if i == numStr {
				count++
			}
		}
		numIStr,_ := strconv.Atoi(string(num[i]))
		if numIStr != count {
			return false 
		}
	}
    
	return true 
}