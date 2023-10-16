package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0,1,2,3,4}, []int{0,1,2,2,1}))
}

func createTargetArray(nums []int, index []int)[]int{
	var output [5]int

    for i:=0;i<len(nums);i++{
		for j:=0;j<len(index);j++{
			output[j] = nums[i]
		}
	}

	return output[:]
}