package main

import "fmt"

func main() {
	nums := []int{0,1,2,3,4}
	index := []int{0,1,2,2,1}
	fmt.Println(createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	target :=make(map[int]int)
	for _,v := range index {
		for _,l := range nums{
			if v == l{
				target[v] = l
			} else if v!=l {
				target[v] = l
			} 
		}
	}
	fmt.Println(target)

    return []int{23}
}