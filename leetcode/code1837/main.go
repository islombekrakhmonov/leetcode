package main

import "fmt"

func main() {
	fmt.Println(sumBase(34,6))
}

func sumBase(n int, k int) int {
	output,remainder := 0,0
	number := n
    for number !=0{
		remainder = number % k
		number /=k
		output += remainder
	}
	return output
}