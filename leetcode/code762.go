package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
}

func countPrimeSetBits(left int, right int) int {
	var output int
	for left <= right {
		if isPrime(countSetBits(left)) {
			output++
		}
		left++
	}
	return output
}

func isPrime(num int) bool{
	divisor :=0
	for i:=2; i<=num; i++{
		if num % i == 0 {
			divisor++
		}
	}
	return divisor==1
}

func countSetBits(num int)int{
    bits := 0
    for ; num > 0; num/=2{
        if num % 2 == 1{
            bits++
        } else{
        }
    }
    return bits
}