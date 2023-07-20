package main

import (
	"fmt"

)

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{{0,0,0},{1,1,1}}))
}

func rowAndMaximumOnes(mat [][]int) []int {
	var count int 
	var count1 int 
	
	mapOfOnes:= make(map[int]int)
	for l,v := range mat{
		count = 0
		for i:=0; i<len(v); i++{
			if v[i] == 1{
				count++
			}
		} 
		if count1 < count {
			count1 = count
			mapOfOnes[count] = l
		}
	}
	output := []int{mapOfOnes[count1], count1}
    
	return output
}