package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2},{-1,-1,-2,-3}}))
}

func countNegatives(grid [][]int) int {
	var output int
	for _,v:= range grid {
		for _,k := range v {
			if k < 0 {
				output++
			}
		} 
	}
    return output
}