package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	fmt.Println(minMaxDifference(90))

	dateTimeUTc := "2025-08-19 05:58:11"
	dateTime, _ := time.Parse("2006-01-02 15:04:05", dateTimeUTc)

	timeNow := time.Now().UTC().Add(-5 * time.Hour)
	fmt.Println(dateTime, timeNow)
	if dateTime.After(timeNow) {
		fmt.Println("yes")
	}
}

func minMaxDifference(num int) int {
	numStr := strconv.Itoa(num)

	seenInt := make(map[byte]bool)
	var (
		max = math.MinInt
		min = math.MaxInt
	)

	for i := 0; i < len(numStr); i++ {
		if !seenInt[numStr[i]] {
			numStrCopyRunes0 := []rune(numStr)
			numStrCopyRunes9 := []rune(numStr)
			seenInt[numStr[i]] = true
			if numStr[i] != '0' {
				numStrCopyRunes0[i] = '0'
			}
			if numStr[i] != '9' {
				numStrCopyRunes9[i] = '9'
			}

			for j := i + 1; j < len(numStr); j++ {
				if numStr[i] == numStr[j] {
					numStrCopyRunes9[j] = '9'
					numStrCopyRunes0[j] = '0'
				}
			}
			formattedNum0, _ := strconv.Atoi(string(numStrCopyRunes0))
			formattedNum9, _ := strconv.Atoi(string(numStrCopyRunes9))

			if formattedNum0 < min {
				min = formattedNum0
			}
			if formattedNum9 > max {
				max = formattedNum9
			}
		}
	}

	return max - min
}
