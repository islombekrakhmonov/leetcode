package main

import "fmt"

func main() {
	fmt.Println(maximumStrongPairXor([]int{1, 2, 3, 4, 5}))

}

func maximumStrongPairXor(nums []int) int {

	maxXor := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			dif, min := absoluteDifferenceAndMin(nums[i], nums[j])
			if dif <= min {
				xor := nums[i] ^ nums[j]
				if maxXor < xor {
					maxXor = xor
				}
			}
		}
	}

	return maxXor
}

func absoluteDifferenceAndMin(x, y int) (difference, min int) {
	if x > y {
		return x - y, y
	}

	return y - x, x
}
