package main

import "fmt"

func main() {
	fmt.Println(twoOutOfThree([]int{1,1,3,2}, []int{2,3}, []int{3}))
	
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var arr []int
	var output []int
    
	for i:=0; i<len(nums1); i++{
		for j:=0; j<len(nums2); j++{
			if nums1[i] == nums2[j] {
				arr = append(arr, nums1[i])
			}
		}
	}

	for i:=0; i<len(nums1); i++{
		for j:=0; j<len(nums3); j++{
			if nums1[i] == nums3[j] {
				arr = append(arr, nums1[i])
			}
		}
	}

	for i:=0; i<len(nums2); i++{
		for j:=0; j<len(nums3); j++{
			if nums2[i] == nums3[j] {
				arr = append(arr, nums2[i])
			}
		}
	}


	for i := 0; i < len(arr); i++ {
		isUnique := true
		for j := 0; j < len(output); j++ {
			if arr[i] == output[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			output = append(output, arr[i])
		}
	}	
	return output
}