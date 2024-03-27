package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {

	var k int

	for i, v := range nums {
		if v != val {
			fmt.Println(nums[k], ",", nums[i], " ", k)
			nums[k] = nums[i]
			k++
		}
	}

	fmt.Println(nums)

	return k
}
