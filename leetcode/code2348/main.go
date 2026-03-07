package main

import "fmt"

func main() {
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
}

func zeroFilledSubarray(nums []int) int64 {
	var (
		output int64
		count  int64
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			output += (count + 1) * count / 2
			count = 0
		}
	}

	output += count * (count + 1) / 2

	return output
}

//  ((n+1)*n)/2
