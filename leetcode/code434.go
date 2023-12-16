package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSegments("           "))
}

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	joined := strings.Split(s, " ")
	var count int 
	for _, segment := range joined {
		if segment != "" {
			count++
		}
	}
	return count
}