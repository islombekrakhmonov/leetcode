package main

import (
	"fmt"
)

func main() {
	fmt.Println(reformatNumber("9964-"))
}

func reformatNumber(number string) string {
	trimmed := ""
	for i:=0; i<len(number); i++{
		if number[i] != ' ' && number[i] != '-'{
			trimmed += string(number[i])
		}
	}

	if len(trimmed) <= 3 {
		return trimmed
	} 

	var output string

	if len(trimmed) > 4 {
		for i:=0; i<3; i++{
			output += string(trimmed[i])
		}
		trimmed = trimmed[3:]
		output += "-"
	} 

	for len(trimmed) != 0 {
		if len(trimmed) > 4 {
			for i:=0; i<3; i++{
				output += string(trimmed[i])
			}
			trimmed = trimmed[3:]
			output += "-"
		} else if len(trimmed) <= 3 {
			for i:=0; i<len(trimmed); i++{
				output += string(trimmed[i])
			}
			trimmed = trimmed[len(trimmed):]
			output += "-"
		} else if len(trimmed) == 4 {
			j := 2
			for j != 0 {
				for i:=0; i<2; i++{
					output += string(trimmed[i])
				}
				trimmed = trimmed[2:]
				output += "-"
				j--
			}
		}
	}

	return	output[:len(output)-1]
}
