package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}

func largestGoodInteger(num string) string {

	triples := []string{"999", "888", "777", "666", "555", "444", "333", "222", "111", "000"}
	for i := 0; i < len(triples); i++ {
		if strings.Contains(num, triples[i]) {
			return triples[i]
		}
	}
	
	return ""
}

func insertionSort(arr []string) {
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