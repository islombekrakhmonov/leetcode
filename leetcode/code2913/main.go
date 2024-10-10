package main

import "fmt"

func main() {

	fmt.Println(sumCounts([]int{1, 2, 1}))
}

func sumCounts(nums []int) int {
	var output int
	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			subarray := nums[startIndex:endIndex]
			countMap := make(map[int]int)
			for _, v := range subarray {
				countMap[v]++
			}
			output += len(countMap) * len(countMap)
		}
	}

	return output
}
