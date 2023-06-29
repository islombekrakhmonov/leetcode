package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(9))
}

func isPowerOfThree(n int) bool {
    base := 1
	exponent := 1000
	var result int 

	for i:= 0; i<exponent; i++{
		if i == 0 {
			result = 1
		} else {
			base *= 3
			result = base 
		}
		if result == n {
			return true
		} 
		if result > n {
			return false 
		}
	}
	return false 
}