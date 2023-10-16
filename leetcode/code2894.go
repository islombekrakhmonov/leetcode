package main

import "fmt"

func main() {
	fmt.Println(differenceOfSums(10, 3))
}

func differenceOfSums(n int, m int) int {
	var sum1 int
	var sum2 int
    for i:=1; i<=n; i++{
		if i % m != 0 {
			sum1 += i
		} else {
			sum2 += i
		}
	}
	return sum1 - sum2
}