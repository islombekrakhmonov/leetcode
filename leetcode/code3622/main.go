package main

import "fmt"

func main() {
	fmt.Println(checkDivisibility(99))
}

func checkDivisibility(n int) bool {
	nStr := fmt.Sprint(n)
	sum, product := 0, 1

	for _, digit := range nStr {
		digitInt := int(digit - '0')
		sum += digitInt
		product *= digitInt
	}

	return n%(sum+product) == 0
}
