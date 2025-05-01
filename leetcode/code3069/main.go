package main

import "fmt"

func main() {
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(resultArray([]int{5, 4, 3, 8}))
}

func resultArray(nums []int) []int {
	var arr1, arr2 []int

	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])

	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}

	arr1 = append(arr1, arr2...)

	return arr1
}
