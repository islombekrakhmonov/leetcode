package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(returnMap("Hello World Hello"))
}

func returnMap (s string ) map[string]int{
	output := make(map[string]int)
	sArray := strings.Fields(s)
	
	fmt.Println(sArray)
	for _,v := range sArray{
		output[v]++
	}
	
	return output
}