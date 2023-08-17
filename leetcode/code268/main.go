package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingNumber([]int{0,1,2}))
}



func missingNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}

	if len(nums) == 1 && nums[0] == 1 {
		return 0
	}

	if len(nums) == 2 && nums[0] == 0 && nums[1] == 1 {
		return 2
	}

    for i:=0; i<=len(nums);i++{
		if !isInSlice(nums, i) {
			return i
		}
		
	}
	return 0
}

func isInSlice(slice []int, num int) bool {
    index := sort.Search(len(slice), func(i int) bool {
        return slice[i] >= num
    })
    return index < len(slice) && slice[index] == num
}