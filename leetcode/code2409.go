package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-16"))
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	year := "2023-"
    layout := "2006-01-02"

    startAlice, _ := time.Parse(layout, year+arriveAlice)
    endAlice, _ := time.Parse(layout, year+leaveAlice)
    startBob, _ := time.Parse(layout, year+arriveBob)
    endBob, _ := time.Parse(layout, year+leaveBob)

    if startAlice.After(endBob) || startBob.After(endAlice) {
        return 0
    }

	overlapStart := max(startAlice, startBob)
    overlapEnd := min(endAlice, endBob)

    days := int(overlapEnd.Sub(overlapStart).Hours() / 24) + 1

    return days
}

func max(a, b time.Time) time.Time {
    if a.After(b) {
        return a
    }
    return b
}

func min(a, b time.Time) time.Time {
    if a.Before(b) {
        return a
    }
    return b
}