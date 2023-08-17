package main

import "fmt"

func main() {
	nums := []int{0,1,2,3,4}
	index := []int{0,1,2,2,1}
	fmt.Println(createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	var output []int
	for i, val := range nums {
		if index[i] < len(output) {
            output = append(output[:index[i]], append([]int{val}, output[index[i]:]...)...)
        } else {
			fmt.Println(output)
            output = append(output, val)
        }	
	}

    return output
}