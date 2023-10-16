package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(sortByBits([]int{0,1,2,3,4,5,6,7,8}))
}

func sortByBits(arr []int) []int {
	mapped := make(map[int]int)
	for i:=0;i<len(arr);i++{
		binaryStr := strconv.FormatInt(int64(arr[i]), 2)
		count := 0
		for _,k := range binaryStr{
			if k == '1' {
				count ++
			}
		}
		mapped[arr[i]] = count
	}
	
	sort.Slice(arr, func(i, j int) bool {
		if mapped[arr[i]] == mapped[arr[j]] {
			return arr[i] < arr[j] 
		}
		return mapped[arr[i]] < mapped[arr[j]] 
	})
	
	return arr
}