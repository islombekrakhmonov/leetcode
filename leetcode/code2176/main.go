package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{3,1,2,2,2,1,3}, 2))
}

func countPairs(nums []int, k int) int {
	var output int 

	for i := 0; i<len(nums); i++ {
		for j := i+1; j<len(nums); j++{
			if nums[i] == nums[j] && (i * j)%k==0{
				output ++
			}
		}
	}
	return output
}

//- nums[0] == nums[6], and 0 * 6 == 0, which is divisible by 2.