package main

import "fmt"

func main() {
	// fmt.Println(check([]int{3, 4, 5, 1, 2}))
	// fmt.Println(check([]int{2, 3, 4, 1}))
	fmt.Println(check([]int{6, 10, 6}))
	// fmt.Println(check([]int{1, 2, 3, 4, 5}))
	// fmt.Println(check([]int{1, 2, 3, 4, 5, 6}))
}

func check(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	// Find the rotation index
	rotationIndex := -1
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			rotationIndex = i + 1
			break
		}
	}

	// If no rotation found, it's already sorted
	if rotationIndex == -1 {
		return true
	}

	rotated := rotateLeft(nums, rotationIndex)
	return isSorted(rotated)
}

func rotateLeft(arr []int, x int) []int {
	n := len(arr)
	x = x % n
	return append(arr[x:], arr[:x]...)
}

func isSorted(arr []int) bool {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
