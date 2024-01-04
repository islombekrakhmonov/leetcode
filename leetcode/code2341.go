package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{5,12,53,22,7,59,41,54,71,24,91,74,62,47,20,14,73,11,82,2,15,38,38,20,57,70,86,93,38,75,94,19,36,40,28,6,35,86,38,94,4,90,18,87,24,22}))
}

func numberOfPairs(nums []int) []int {
	var countPairs int
	var leftOver int 
	numbers := make(map[int]int)
	for i := 0; i<len(nums); i++{
		numbers[nums[i]]++
	}

	fmt.Println(numbers)
	for _, c := range numbers{ 
		countPairs += c / 2 
		if c%2 != 0 {
			leftOver++
		}
	}
    
	return []int{countPairs,leftOver}
}
