package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(commonFactors(12,6))
}

func commonFactors(a int, b int) int {
	var max = math.Max(float64(a),float64(b))
	var count = 0;

	for i := 1; i <= int(max); i++{
		
		if a % i == 0 && b % i == 0 {
			count++
		}

	} 
	return count
    
}