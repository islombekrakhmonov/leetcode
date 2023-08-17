package main

import "fmt"

func main() {
	fmt.Println(findFinalValue([]int{5,6,3,1,12}, 3))
}

func findFinalValue(nums []int, original int) int {
	j:=1

	for j > 0 {
		for i:=0; i<len(nums);i++{
			if original == nums[i]{
				original = 2 * original
				j++
			}
		}
		j--
	}
    
	return original
}