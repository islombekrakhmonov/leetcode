package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isFascinating(100))
}

func isFascinating(n int) bool {
	Concatenated1 := strconv.Itoa(n*2)
	Concatenated2 := strconv.Itoa(n*3)
	final := strconv.Itoa(n)
	final += Concatenated1 + Concatenated2
	count := make(map[int]int)

	for i:=1;i<=len(final);i++{
		for j:=0;j<len(final); j++{
			if strconv.Itoa(i) == string(final[j]) {
				count[i]++
			}
		}
	}

	for _,v := range count{
		if v != 1 {
			return false 
		}
	}

	return len(count) == 9
}