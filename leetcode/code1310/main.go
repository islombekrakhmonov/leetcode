package main

import "fmt"

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

	fmt.Println(prefix)
	result := make([]int, len(queries))
	for i, q := range queries {
		left, right := q[0], q[1]
		result[i] = prefix[right+1] ^ prefix[left]
	}

	return result
}
