package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1,2,3},[]int{3,2,7}, 4))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var output int

	for i:=0; i<len(startTime); i++{
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			output++
		}
	}
    
	return output
}