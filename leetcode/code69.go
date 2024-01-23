package main

import "fmt"

func main() {
	
	fmt.Println(mySqrt(8))

}

func mySqrt(x int) int {

	if x < 0 {
		return -1
	}


	low, high := 0, x

	for low <= high {
		mid := (low + high)/ 2
		guess := mid * mid 
		if guess == x {
			return guess
		} else if guess > x {
			high = mid - 1 
			fmt.Println(high, guess)
		} else {
			fmt.Println(low, mid, guess)
			low = mid + 1 
		}
	}

	return low -1 
}
