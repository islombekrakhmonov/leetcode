package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {	
	first, second := 0, len(numbers)-1

	for first < second {
		if numbers[first]+numbers[second] == target && first<second{
			return []int{first+1, second+1}
		} else if numbers[first]+numbers[second] < target {
			first++
		} else {
			second--
		}
	}

	return nil
}