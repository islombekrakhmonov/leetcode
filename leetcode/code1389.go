package main

import "fmt"

func main() {
	nums := []int{0,1,2,3,4}
	index := []int{0,1,2,2,1}
	fmt.Println(createTargetArray(nums,index))
}

func createTargetArray(nums []int, index []int) []int {
	gross := make(map[int]int)
    var output []int
	for _,v := range index{
		for _,l := range nums{
			if v == l {
				gross[v] = l
			} else if v!=l {
				gross[v] = l
			} 
		}
	}
	fmt.Println(gross)


	return output


}