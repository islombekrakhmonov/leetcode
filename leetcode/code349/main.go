package main

import "fmt"

func main() {

	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))

}

func intersection(nums1 []int, nums2 []int) []int {
	var output []int
	hashSet := make(map[int]int)
	hashSetNums1 := make(map[int]int)

	for _, v := range nums1 {
		hashSetNums1[v]++
	}

	for _, k := range nums2 {
		if _, exists := hashSetNums1[k]; exists {
			hashSet[k]++
		}
	}

	for i := range hashSet {
		output = append(output, i)
	}

	return output
}
