package main

import (
	"fmt"
)

func main() {
	fmt.Println(distanceBetweenBusStops([]int{3,6,7,2,9,10,7,16,11}, 6,2))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	l := len(distance)
	minDistance := 0
	minDistanceBackwards := 0

    for i:=start;i!=destination; i = (i + 1) % l{
		minDistance += distance[i]
	}

	for i:=destination; i!=start;i = (i + 1) % l{
		minDistanceBackwards += distance[i]
	}
			
	if minDistanceBackwards > minDistance {
		return minDistance
	} else {
		return minDistanceBackwards
	}
}