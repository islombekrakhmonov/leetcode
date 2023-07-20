package main

import "fmt"

func main() {
	fmt.Println(maximumCount([]int{-3,-2,-1,0,0,1,2}))
}

func maximumCount(nums []int) int {
	max,pos, neg := 0,0,0


	for _,v := range nums{
		if v >0 {
			pos++
		} else if v < 0{
			neg++
		}
	}
	if pos>neg{
		max = pos
	} else {
		max = neg
	}
    
	return max
}