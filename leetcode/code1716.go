package main

import "fmt"

func main() {
	fmt.Println(totalMoney(10))
}

func totalMoney(n int) int {
	var total int

	monday := 1
	day := 0

	for i := 1; i <= n; i++ {
		total += monday + day

		day++

		if day == 7 {
			day = 0
			monday++
		}
	}

	return total
}