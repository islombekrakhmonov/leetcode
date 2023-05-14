package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	var result string
	var sample string
	// pop := strings.Split(s, " ")
	// t := strings.Join(slice, " ")
	for _,v := range s {
		sample = string(v) + sample
	  }
	  t := strings.Split(sample, " ")
	  fmt.Println(t)
	  for i:=len(t)-1;i>=0;i--{
		if i==len(t)-1 {
			result = result + t[i]
		} else {
		result = result + " "+t[i]
		} 
		}
	return result
}
