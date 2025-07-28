package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(winningPlayerCount(4, [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}}))

	startDateString := "12.07.2025 13:50"

	fmt.Println(startDateString)

	startDate, err := time.Parse("02.01.2006 15:04", startDateString)
	if err != nil {
		fmt.Println(err)
	}

	formattedDate := startDate.Format("2006-01-02T15:04:05")

	fmt.Println(formattedDate)

	fmt.Println(startDate.Format("02.01.2006 15:04"))

}

func winningPlayerCount(n int, pick [][]int) int {
	numOfWins := 0

	var hashMap = make(map[int][]int)

	for i := 0; i < len(pick); i++ {
		hashMap[pick[i][0]] = append(hashMap[pick[i][0]], pick[i][1])
	}

	for key, array := range hashMap {
		hashColor := make(map[int]int)
		for i := 0; i < len(array); i++ {
			hashColor[array[i]]++
			if hashColor[array[i]] > key {
				numOfWins++
				break
			}
		}
	}

	return numOfWins
}
