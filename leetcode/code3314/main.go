package main

import "fmt"

func main() {
	fmt.Println(minBitwiseArray([]int{2, 3, 5, 7}))
}

func minBitwiseArray(nums []int) []int {
	output := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var j int
		for j <= 1000 {
			if j|(j+1) == nums[i] {
				output[i] = j
				break
			}
			j++
		}
		if output[i] == 0 {
			output[i] = -1
		}
	}

	return output
}
