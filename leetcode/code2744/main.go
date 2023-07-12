package main

import "fmt"

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd","ac","dc","ca","zz"}))
}

func maximumNumberOfStringPairs(words []string) int {
	var reversed []string
	var output int
	for _,v := range words{
		reversedStr := ""
        for j:=len(v)-1; j>=0; j--{
			reversedStr +=string(v[j])
		}
		reversed = append(reversed, reversedStr)
    }
    for i,v := range words{
		
        for j:=i+1; j<len(reversed); j++{
            if v == reversed[j]{
				output++
			}
        }
    }
	return output
}