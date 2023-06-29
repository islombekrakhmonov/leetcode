package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(8))

	var arr [5]string
	fmt.Println(arr)
}

func isPowerOfTwo(n int) bool {
	base := 1
	exponent := 1000
	var result int

	for i:=0; i<exponent; i++{
		if i == 0{ 
			result =1
		} else {
			base *=2
			result = base
		}
		fmt.Println(result)
		if result == n {
			return true
		}
		if result > n {
			return false 
		}
	}
	return false 
}

func reorderedPowerOf2(n int) bool {
    base := 1
	exponent := 1000
	var result int

	for i:=0; i<exponent; i++{
		if i == 0{ 
			result =1
		} else {
			base *=2
			result = base
		}
		fmt.Println(result)
		if result == n {
			return true
		}
		if result > n {
			return false 
		}
	}
	return false 
}