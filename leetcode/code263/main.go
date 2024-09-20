package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}

	// we repeatedly divide it by 2, 3, or 5 if it is divisible by any of these numbers. If we can reduce the number to 1 using these divisions, then it is an ugly number.

	return true
}
