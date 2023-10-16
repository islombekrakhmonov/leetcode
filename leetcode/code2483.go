package main

import "fmt"

func main() {
	fmt.Println(bestClosingTime("YN"))	
}

func bestClosingTime(customers string) int {
    var penalty []int 
	var penaltySum int
	maxSum := 0
	maxSumIndex := -1

	for i:=0;i<len(customers); i++{
		if string(customers[i]) == "Y" {
			penalty = append(penalty, 1)	
		} else {
			penalty = append(penalty, -1)
		}
	}

	for i := 0; i<len(penalty);i++{
		penaltySum += penalty[i]
		if penaltySum > maxSum {
			maxSum = penaltySum
			maxSumIndex = i+1
		}
	}

	if maxSumIndex < 0 {
		return 0
	} else {
		return maxSumIndex
	}
}