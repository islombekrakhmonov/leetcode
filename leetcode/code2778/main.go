package main

import "fmt"

func main() {
	fmt.Println(sumOfSquares([]int{1,2,3,4}))
}

func getElement(slice []int, index int) int {
    i := index - 1

    if i >= 0 && i < len(slice) {
        return slice[i]
    }

    return 0
}

func sumOfSquares(nums []int) int {
	n := len(nums)
	var output int

    for i:=0;i<n;i++{
		if n % (i+1) == 0 {
			output += nums[i] * nums[i]
		}
	}

	return output
}