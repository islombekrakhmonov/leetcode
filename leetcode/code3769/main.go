package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(sortByReflection([]int{8, 2}))
}

type Pair struct {
	num       int
	reflected int
}

func sortByReflection(nums []int) []int {
	pairs := []Pair{}

	for _, num := range nums {
		binaryStr := strconv.FormatInt(int64(num), 2)
		reversed := reverseSlice([]rune(binaryStr))

		reflected, _ := strconv.ParseInt(string(reversed), 2, 64)
		fmt.Println(binaryStr, reversed, reflected)

		pairs = append(pairs, Pair{
			num:       num,
			reflected: int(reflected),
		})
	}

	sort.Slice(
		pairs,
		func(i, j int) bool {
			if pairs[i].reflected == pairs[j].reflected {
				return pairs[i].num < pairs[j].num
			}
			return pairs[i].reflected < pairs[j].reflected
		},
	)

	for i := 0; i < len(nums); i++ {
		nums[i] = pairs[i].num
	}

	return nums
}

func reverseSlice(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
