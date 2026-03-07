package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	var maxLength int

	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			if endIndex-startIndex < 2 {
				continue
			}
			subArray := nums[startIndex:endIndex]
			max := subArray[0]
			min := subArray[0]
			for i := 0; i < len(subArray); i++ {
				if subArray[i] > max {
					max = subArray[i]
				}
				if subArray[i] < min {
					min = subArray[i]
				}
			}
			if max-min == 1 {
				if len(subArray) > maxLength {
					maxLength = len(subArray)
				}
			}
			fmt.Println(maxLength, subArray)
		}
	}

	return maxLength
}
