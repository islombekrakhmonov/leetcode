package main

import (
	"fmt"

)

func main() {
	input := [][]string {{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}
	fmt.Println(input)
	ruleKey:= "type"
	ruleValue := "phone"
	fmt.Println(countMatches(input,ruleKey,ruleValue))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int
	for _,v := range items{
		if ruleKey == "type"{
			if v[0] == ruleValue{
				count++
			}
		} else if ruleKey == "color"{
			if v[1] == ruleValue{
				count++
			}
			
		} else if ruleKey == "name"{
			if v[2] == ruleValue{
				count++
			}
		}
	}
	return count 
}