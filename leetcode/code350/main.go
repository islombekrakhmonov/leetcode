package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {

	hashMap := make(map[int]int)
	var largerArray []int
	var output []int

	// hashMap2 := make(map[int]int)

	if len(nums1) > len(nums2) {
		for _, v := range nums2 {
			hashMap[v]++
		}
		largerArray = nums1
	} else {
		for _, v := range nums1 {
			hashMap[v]++
		}
		largerArray = nums2
	}

	for _, v := range largerArray {
		if _, exists := hashMap[v]; exists {
			hashMap[v]--
			if hashMap[v] == 0 {
				delete(hashMap, v)
			}
			output = append(output, v)
		}
	}

	return output
}

//Iterate over the larger array and check whether its number is contained in the map
// If the number is contained in the map, reduce count to number
// by 1, if count after reduction = 0, remove the number from the map
// We add the found number to the result array
