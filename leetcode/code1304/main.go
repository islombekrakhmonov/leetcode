package main

import (
	"fmt"
)

func main() {
	
	fmt.Println(sumZero(2))
}


func sumZero(n int) []int {
	output := make([]int, n)

	if n%2 == 1 {
		output[n-1] = 0
	}
	for i := 0; i < n/2; i++ {
		output[i] = (i + 1)
		output[n-i-1] = -(i + 1)
	}
    
	return output
}
