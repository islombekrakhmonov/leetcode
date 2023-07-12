package main

import "fmt"

func main() {
	fmt.Println(consecutiveNumbersSum(15))
}

func consecutiveNumbersSum(n int) int {
	var sum int

	for i:=1; i<n/2+1; i++{ 
		for j:=i; j<=n/2+2; j++ {
			if i+j == n {
				fmt.Println(i,j)
				sum++
			}
			if i == 1 && j == n-1 {
				sum--
			}
		}
	}
	return sum+1
}