package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int {

	var pair int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			division := nums2[j] * k
			if nums1[i] % division == 0 {
				pair++
			}
		}
	}

	return pair
}
