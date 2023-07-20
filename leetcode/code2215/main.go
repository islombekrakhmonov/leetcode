package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{-68,-80,-19,-94,82,21,-43}, []int{-63}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var answer [2][]int
	mapped1 := make(map[int]int)
	mapped2 := make(map[int]int)


    for _,v:= range nums1{
		count := 0
		for _,j := range nums2{
			if v != j {
				count ++
			}
		}
		if count == len(nums2) {
			if _, ok := mapped1[v]; !ok {
				mapped1[v] = v
				answer[0] = append(answer[0], mapped1[v])
			}
		}
	}

	for _,v:= range nums2{
		count := 0
		for _,j := range nums1{
			if v != j {
				count ++
			}
		}
		if count == len(nums2) {
			if _, ok := mapped1[v]; !ok {
				mapped2[v] = v
				answer[1] = append(answer[1], mapped2[v])
			}
		}
	}
	return answer[:]
}