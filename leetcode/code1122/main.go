package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var output []int
	var notInArr2 []int

	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				output = append(output, arr1[j])
			}
		}
	}

	for _, v := range arr1 {
		if !contains(v, arr2) {
			notInArr2 = append(notInArr2, v)
		}
	}

	if len(notInArr2) > 0 {
		sort.Ints(notInArr2)
		output = append(output, notInArr2...)
	}

	return output
}

func contains(num int, slice []int) bool {
	for _, v := range slice {
		if num == v {
			return true
		}
	}

	return false
}
