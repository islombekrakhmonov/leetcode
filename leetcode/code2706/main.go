package main

import (
	"fmt"
)

func main() {
	fmt.Println(buyChoco([]int{3, 2, 3}, 3))

}

func buyChoco(prices []int, money int) int {

	insertionSort(prices)

	if prices[0]+prices[1] <= money {
		return money - prices[0] - prices[1]
	}

	return money
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
