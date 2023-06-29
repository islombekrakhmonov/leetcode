package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	var output []int
    last := digits[len(digits)-1]
	lastPlusOne := last + 1
	if len(digits) < 2 && last == 9 {
		output = append(output, 1)
		output = append(output, 0)
	} else if len(digits)>1 && last == 9 && digits[len(digits)-2] == 9 && len(digits)<3{
		output = digits[:len(digits)-2]
		output = append(output, 1)
		output = append(output, 0)
		output = append(output, 0)
	} else if len(digits)>2 && last == 9 && digits[len(digits)-2] == 9 && digits[len(digits)-3] == 9 && len(digits)<4{
		output = digits[:len(digits)-3]
		output = append(output, 1)
		output = append(output, 0)
		output = append(output, 0)
		output = append(output, 0)
	} else if len(digits)>3 && last == 9 && digits[len(digits)-2] == 9 && digits[len(digits)-3] == 9 && digits[len(digits)-4] == 9{
		output = digits[:len(digits)-4]
		output = append(output, 1)
		output = append(output, 0)
		output = append(output, 0)
		output = append(output, 0)
		output = append(output, 0)
	} else {
		output = digits[:len(digits)-1]
		if last == 10 {
			output = append(output, 1)
			output = append(output, 0)
		} else {
			output = append(output, lastPlusOne)
		}
	}
	
	return output
}