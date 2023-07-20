package main

import "fmt"

func main() {
	fmt.Println(heightChecker([]int{1,1,4,2,1,3}))
}

func heightChecker(heights []int) int {
	var output int
	var notSorted []int
	l := len(heights)
	notSorted = append(notSorted, heights...)	

	for i:=0;i<l; i++{
		for j:=0; j<l-i-1;j++{
			if heights[j] > heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
			}
		}
	}

	for i:=0;i<l;i++{
		if notSorted[i] != heights[i]{
			output++
		}
	}

	return output
}