package main

import "fmt"

func main() {
	fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
}

func hardestWorker(n int, logs [][]int) int {

	startTime := 0
	maxWorkingTime := 0
	maxWorkingTimeWorker := 0

	for i := 0; i < len(logs); i++ {
		if i != 0 {
			startTime = logs[i-1][1]
		}
		workingTime := logs[i][1] - startTime
		if workingTime > maxWorkingTime {
			maxWorkingTime = workingTime
			maxWorkingTimeWorker = logs[i][0]
		} else if workingTime == maxWorkingTime {
			if logs[i][0] < maxWorkingTimeWorker {
				maxWorkingTimeWorker = logs[i][0]
			}
		}
	}

	return 0
}
