package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var answer [2][]int
	mapped1 := make(map[int]bool)
	mapped2 := make(map[int]bool)

	for _, v := range nums1 {
		mapped1[v] = false
	}

	for _, v := range nums2 {
		mapped2[v] = false
	}

	for _, v := range nums1 {
		if _, exists := mapped2[v]; !exists {
			answer[0] = append(answer[0], v)
			mapped2[v] = false
		}
	}

	for _, v := range nums2 {
		if _, exists := mapped1[v]; !exists {
			answer[1] = append(answer[1], v)
			mapped1[v] = false
		}
	}

	return answer[:]
}
