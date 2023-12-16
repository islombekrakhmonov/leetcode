package main

import "fmt"

func main() {
	fmt.Println(findIntersectionValues([]int{4,3,2,3,1}, []int{2,2,5,2,3,6}))
}
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	count1, count2  := 0,0
	for i :=0; i<len(nums1); i++ {
		for j:=0; j<len(nums2); j++{
			if nums1[i] == nums2[j] && i < len(nums1) {
				count1++
				break
			}
		}
	}

	for i :=0; i<len(nums2); i++ {
		for j:=0; j<len(nums1); j++{
			if nums2[i] == nums1[j] && i < len(nums2) {
				count2++
				break
			}
		}
	}

	return []int{count1, count2}
}


