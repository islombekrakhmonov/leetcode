package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

func lastStoneWeight(stones []int) int {

	k := len(stones)
	for k > 0 {
		max1, max2 := 0, 0
		max1Index, max2Index := 0, 0
		for i, v := range stones {
			if max1 <= v {
				max2 = max1
				max1 = v
				max1Index = i
			} else if v > max2 && max1 != v {
				max2 = v
				max2Index = i
			}
		}
		newSlice := append(stones[:max2Index], stones[max2Index:]...)
		fmt.Println(max1Index, max2Index)
		k--
		fmt.Println(newSlice)
		if max1 == max2 {
		}

	}

	return 0
}
