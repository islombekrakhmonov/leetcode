package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{5, 6, 7, 8, 9}, 9))
	fmt.Println(binarySearch([]int{2, 3, 4, 7, 11}, 11))
}

func findKthPositive(arr []int, k int) int {
	max := arr[len(arr)-1]
	var missing []int
	seenMap := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		seenMap[arr[i]] = true
	}

	for i := 1; i < max; i++ {
		if !seenMap[i] {
			missing = append(missing, i)
		}
	}
	for len(missing) < k {
		max++
		missing = append(missing, max)
	}

	return missing[k-1]
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
