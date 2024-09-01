package main

import "fmt"

func main() {
	fmt.Println(findPeaks([]int{2, 4, 4}))
}

func findPeaks(mountain []int) []int {
	peaks := make([]int, 0)

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			peaks = append(peaks, i)
		}
	}
	return peaks
}
