package main

import "fmt"

func main() {
	fmt.Println(arrayPairSum([]int{5,2,7,1}))
}

func arrayPairSum(nums []int) int {
	var output int 
	n := len(nums)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
		n--
	}

	for i := 0; i<len(nums); i += 2 {
		if nums[i] > nums[i+1] {
			output += nums[i+1]
		} else {
			output += nums[i]
		}
	}
    
	return output
}