package main

import "fmt"

func main() {
	sortColors([]int{1,2,0})
}

func sortColors(nums []int) {

	left, right := 0, len(nums)-1

	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			fmt.Println(nums, i)
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			fmt.Println(nums, i)
		} else {
			fmt.Println(nums, i)
			i++
		}
	}
}
