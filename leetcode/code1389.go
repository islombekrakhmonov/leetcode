package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0,1,2,3,4}, []int{0,1,2,2,1}))
}

func createTargetArray(nums []int, index []int)[]int{
	target := make([]int, len(nums))

	for idx, i := range index {
        for j := len(target)-1; j > i; j-- {
            target[j] = target[j-1]
        }
        
        target[i] = nums[idx]
    }

	return target
}