package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{3,6,1,0}))
}

func dominantIndex(nums []int) int {
	var output [2]int 
	for i,v := range nums{
		if v > output[0] {
			output[0] = v
			output[1] = i
		}
	}

	for _,v := range nums {
		if output[0] != v && v != 0 && output[0]/v < 2 {
			return -1 
		}
	}
    
	return output[1]
}