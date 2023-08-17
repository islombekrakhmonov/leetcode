package main

import "fmt"

func main() {
	fmt.Println(rearrangeArray([]int{3,1,-2,-5,2,-4}))
}

func rearrangeArray(nums []int) []int {
	positives, negatives := []int{}, []int{}
	var output []int
	for _,v := range nums{
		if v >0 {
			positives = append(positives, v)
		} else {
			negatives = append(negatives, v)
		}
	}

	for i := 0; i < len(positives) && i < len(negatives); i++ {
		output = append(output, positives[i])
		output = append(output, negatives[i])
	}
	return output
}