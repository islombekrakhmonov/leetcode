package main

import "fmt"

func main() {
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}))
}

func isArraySpecial(nums []int) bool {

	for i := 1; i < len(nums); i++ {
		num1 := nums[i]
		num2 := nums[i-1]
		if (num1%2 == 0 && num2%2 == 0) || (num1%2 != 0 && num2%2 != 0) {
			return false
		}
	}

	return true
}
