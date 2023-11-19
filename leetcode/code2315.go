package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAsterisks("*||*|||||*|*|***||*||***|"))
}

func countAsterisks(s string) int {
	var output int
	var indexes []int
	for i:=0; i<len(s); i++{
		for j:=i+1; j<len(s); j++{
			if s[i] == '|' && s[j] == '|' {
				for k:=i; k<=j; k++{
					indexes = append(indexes, k)
				}
				fmt.Println(i , " aSD", j)
				i = j
				break
			}
		}
	}

	fmt.Println(indexes)

	for i:=0; i<len(s); i++{
		found := false
		for _, v := range indexes {
			if i == v {
				found = true
				break
			}
		}
		if !found && s[i] == '*' {
			output++
		}
	} 
	return output
}