package main

import "fmt"

func main() {
	fmt.Println(minimumChairs("ELELEEL"))
}

func minimumChairs(s string) int {

	var count int
	var maxCount int
	for _, v := range s {
		if v == 'E' {
			count++

		} else if v == 'L' {
			count--
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
