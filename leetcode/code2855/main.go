package main

import "fmt"

func main() {
	fmt.Println(minimumRightShifts([]int{1, 2, 3, 4, 5}))
}

func minimumRightShifts(nums []int) int {
	var minShifts int

	for i := 0; i < len(nums); i++ {
		if isAsceding(nums) {
			return minShifts
		}
		nums = rightShift(nums)
		minShifts++
	}

	return -1
}

func rightShift(arr []int) []int {
	shift := 1

	shift = shift % len(arr)

	return append(arr[len(arr)-shift:], arr[:len(arr)-shift]...)
}

func isAsceding(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}
