package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(twoSum2([]int{3,2,4}, 6))
}


func twoSum(nums []int, target int) []int {
	var output []int
	for k,v := range nums{
		for i:=k+1;i<len(nums); i++{
			if v+nums[i] == target{
				output = append(output, k)
				output = append(output, i)
				return output
			}
		}
	}
	 
	return output
}

func twoSum1(nums []int, target int) []int {
	numIndices := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, found := numIndices[complement]; found {
            return []int{index, i}
        }
        numIndices[num] = i
		fmt.Println(numIndices)
    }

    return nil
 }
 

 func twoSum2(nums []int, target int) []int {

	sort.Ints(nums)
	first, second := 0, len(nums)-1
	
	for first < second {
		if nums[first]+nums[second] == target {
			return []int{first, second}
		} else if nums[first]+nums[second] < target {
			first++
		} else {
			second--
		}
	}

	return nil
 }
