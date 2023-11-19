package main

import "fmt"

func main() {
	fmt.Println(pivotInteger(1))
}

func pivotInteger(n int) int {
	var num1 int

    for i:=n; i>0; i-- {
		num1 += i
		num2 := 0
		for j:=1; j<=i; j++{
			num2 += j
		}
		if num1 == num2 {
			return i
		}
	}

	return -1
}