package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{3, 7, 6, 4, 3, 0, 1, 0}))
}

func validMountainArray(arr []int) bool {
	isDecreasing := false
	isIncreasing := false
	increased := false
	decreased := false

	if len(arr) < 3 {
		return false
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			isDecreasing = true
			decreased = true
			isIncreasing = false
		} else if arr[i] < arr[i+1] {
			isIncreasing = true
			increased = true
			isDecreasing = false
			if decreased {
				return false
			}
		} else {
			return false
		}
	}
	if !isDecreasing || isIncreasing || !increased || !decreased {
		return false
	}

	return true
}
