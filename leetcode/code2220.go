package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minBitFlips(10,7))
	
}

func minBitFlips(start int, goal int) int {
	var output int
	maxLen := maxBinaryLength(start, goal)

   	binaryStart := formatBinary(start, maxLen)
	binaryGoal := formatBinary(goal, maxLen)

	for i,v := range binaryStart {
		for j,k := range binaryGoal	{
			if v != k && i == j{
				output++
				break
			} 
		}
	}

	return output
}

func maxBinaryLength(nums ...int) int {
	maxLen := 0
	for _, num := range nums {
		binaryLen := len(strconv.FormatInt(int64(num), 2))
		if binaryLen > maxLen {
			maxLen = binaryLen
		}
	}
	return maxLen
}

func formatBinary(num, length int) string {
	binaryStr := strconv.FormatInt(int64(num), 2)
	for len(binaryStr) < length {
		binaryStr = "0" + binaryStr
	}
	return binaryStr
}