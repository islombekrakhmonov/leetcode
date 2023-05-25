package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(10))
}

func sumOfMultiples(n int) int {
	var slc []int
	var sum int
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			slc = append(slc, i)
		}
	}
	for i := 0; i < len(slc); i++ {
		sum += slc[i]
	}
	return sum
}

