package main

import "fmt"

func main() {
	fmt.Println(countOdds(14,17))
}

func countOdds(low int, high int) int {	
	if low % 2 != 0 {
		return ((high - low) / 2 ) +1
	} else {
		return ((high - low) / 2 ) 
	}
}