package main

import "fmt"

func main() {
	fmt.Println(gcdOfOddEvenSums(4))
}

func gcdOfOddEvenSums(n int) int {
	sumOdd, sumEven := 0, 0
	for i := 1; n*2+1 > i; i++ {
		if i%2 == 0 {
			sumEven += i
		} else {
			sumOdd += i
		}
	}

	return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
