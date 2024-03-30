package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	var output []string

	if len(nums) == 0 {
		return output
	}

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// If the current number is consecutive to the previous number, update the end point of the range
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			// If not consecutive, add the range to the output list
			output = append(output, formatRange(start, end))
			// Update the start and end points of the new range
			start, end = nums[i], nums[i]
		}
	}

	output = append(output, formatRange(start, end))

	return output
}

func formatRange(start, end int) string {
	// If the start and end points are the same, return a single number
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	// Otherwise, return a range
	return fmt.Sprintf("%d->%d", start, end)
}
