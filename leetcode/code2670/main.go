package main

import "fmt"

func main() {
	fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
}

func distinctDifferenceArray(nums []int) []int {

	diff := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		prefix := nums[0 : i+1]
		suffix := nums[i+1:]
		countPrefix := make(map[int]struct{})
		countSuffix := make(map[int]struct{})
		for _, v := range prefix {
			countPrefix[v] = struct{}{}
		}

		for _, v := range suffix {
			countSuffix[v] = struct{}{}
		}

		diff[i] = len(countPrefix) - len(countSuffix)
	}

	return diff

}
