package main

import (
	"fmt"
)

func main() {
	fmt.Println(intersection([][]int{{44, 21, 48}}))
}

func intersection(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	elementCount := make(map[int]int)

	for _, v := range nums[0] {
		elementCount[v] = 1 // Initially set count to 1 for elements in the first array
	}

	for i := 1; i < len(nums); i++ {
		currentArray := make(map[int]bool) // To avoid counting duplicates within the same array
		for _, num := range nums[i] {
			if _, exists := elementCount[num]; exists && !currentArray[num] {
				elementCount[num]++
				currentArray[num] = true
			}
		}
	}

	var result []int
	for num, count := range elementCount {
		if count == len(nums) { // If an element appears in all arrays
			result = append(result, num)
		}
	}

	insertionSort(result)

	return result
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
