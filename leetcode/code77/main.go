package main

import "fmt"

func main() {
	fmt.Println(combine(4,2))
}

func combine(n int, k int) (output [][]int) {
	arr, i := make([]int, k), 0
	for i >= 0 {	
		arr[i]++
		fmt.Println(arr)
		if arr[i] > n {
			i--
		} else if i == k - 1 {
			output = append(output, append([]int{}, arr...))
		} else {
			i++
			arr[i] = arr[i - 1]
		}
	}
	return output
}

