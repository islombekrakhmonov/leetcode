package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(findDelayedArrivalTime(25, 5))
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {

	midnight := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

	resultTime := midnight.Add(time.Duration(arrivalTime+delayedTime) * time.Hour)
    
	return resultTime.Hour()	
}