package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}))
	fmt.Println(nextGreaterElement1([]int{4,1,2}, []int{1,3,4,2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var output []int

	for i:=0; i<len(nums1); i++ {
		var found bool
		for j:=0; j<len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for k:=j; k<len(nums2); k++ {
					if nums2[j] < nums2[k]{
						output = append(output, nums2[k])
						found = true
						break
					}
				}
			}
		}
		if !found {
			output = append(output, -1)
		}
	}

	return output
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
    dict := make(map[int]int)
    stack := []int{}

    for _, num := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < num {
            removed := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            dict[removed] = num
			fmt.Println(dict)
        }
        stack = append(stack, num)
    }
	fmt.Println(stack)

    res := []int{}
    for _, num := range nums1 {
        if value, ok := dict[num]; ok {
            res = append(res, value)
        } else {
            res = append(res, -1)
        }
    }
    return res
}