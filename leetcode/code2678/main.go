package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522","5303914400F9211","9273338290F4010"}))
}

func countSeniors(details []string) int {
    var output int

	for _, v:= range details {
		age, _ := strconv.Atoi(v[11:13]) 
		
		if age >= 60 {
			output ++
		}
	}
	
	return output
}