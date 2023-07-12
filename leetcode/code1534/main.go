package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countGoodTriplets([]int{3,0,1,1,9,7}, 7,2,3))
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var output int 
	for i:=0; i<len(arr); i++{
		for j:=i+1; j<len(arr); j++{
			for k:=j+1; k<len(arr); k++{
				if int(math.Abs(float64(arr[i] - arr[j]))) <= a {
					if int(math.Abs(float64(arr[j] - arr[k]))) <= b {
						if int(math.Abs(float64(arr[i] - arr[k]))) <= c {
							output++
						}
					}
				}
			}
		} 
	}
	return output
}
