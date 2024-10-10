package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}

func findSpecialInteger(arr []int) int {

	l := len(arr)

	countMap := make(map[int]int)

	for i := 0; i < l; i++ {
		countMap[arr[i]]++
		if countMap[arr[i]] > (l / 4) {
			return arr[i]
		}
	}

	return 0
}
