package main

import "fmt"

func main() {
	fmt.Println(smallestEvenMultiple(3))
}

func smallestEvenMultiple(n int) int {
	for true{
		if n % 2 == 0 {
			return n
		} else {
			n *= 2
		} 
		break 
	}
	return n
}