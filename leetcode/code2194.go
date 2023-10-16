package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(cellsInRange("K1:L2"))
}

func cellsInRange(s string) []string {
	var output []string
    rowLength, _ := strconv.Atoi(string(s[len(s)-1]))
	rowStart, _ := strconv.Atoi(string(s[1]))
	
	for char :=  s[0]; char <= s[3]; char++ {
		for i:=rowStart; i<=rowLength; i++{
			iStr := strconv.Itoa(i)
			finalStr := string(char) + iStr
			output = append(output, finalStr)
		}
	}

	return output
}