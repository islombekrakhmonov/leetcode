package main

import "fmt"

func main() {
	fmt.Println(countOperations(6,4))
}

func countOperations(num1 int, num2 int) int {
    var opertation int 
	for num1 != 0 && num2 != 0{
		if num1 >= num2 {
			fmt.Println(num1)
			num1 -= num2 
			fmt.Println(num1)
			opertation ++
		} else {
			fmt.Println(num1)
			num2 -= num1
			fmt.Println(num2)
			opertation ++
		}
	}
	return opertation
}