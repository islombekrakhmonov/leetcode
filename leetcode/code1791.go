package main

import "fmt"

func main() {
	fmt.Println(findCenter([][]int{{1,2},{2,3},{4,2}}))
}

func findCenter(edges [][]int) int {
	var output, maxFreq int

	numbers := make(map[int]int)
	for _, array := range edges {
		for i:=0; i<len(array); i++{
			numbers[array[i]]++
		}
	}
	
	for num, freq := range numbers {
		if freq > maxFreq {
			output = num
            maxFreq = freq
		}
	}

	return output
}