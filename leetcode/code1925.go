package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countTriples(10))
}

func countTriples(n int) int {
	var count int 

	for a := 1; a<=n; a++{
		for b:=1; b<=n; b++ {
			c := math.Sqrt(float64(a*a) + float64(b*b))
			if c == float64(int(c)) && int(c) <= n {
				count++
			}
		}
	} 
    
	return count
}