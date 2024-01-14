package main

import (
	"fmt"
)

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
}

func garbageCollection(garbage []string, travel []int) int {

	collectionTime := 0
	gTravelTime := 0
	pTravelTime := 0
	mTravelTime := 0
	currCum := 0

	for i, gCollect := range garbage {
		if i != 0 {
			currCum += travel[i-1]
		}
		collectionTime += len(gCollect)
		for j := 0; j < len(gCollect); j++ {
			if gCollect[j] == 'M' {
				mTravelTime = currCum
			} else if gCollect[j] == 'P' {
				pTravelTime = currCum
			} else if gCollect[j] == 'G' {
				gTravelTime = currCum
			}
		}
	}

	return collectionTime + gTravelTime + pTravelTime + mTravelTime
}
