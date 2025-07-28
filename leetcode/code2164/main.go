package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortEvenOdd([]int{4, 1, 2, 3}))

}

func sortEvenOdd(nums []int) []int {

	var (
		odd  []int
		even []int
	)

	for i, v := range nums {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sort.Ints(even)

	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = even[0]
			even = even[1:]
		} else {
			nums[i] = odd[0]
			odd = odd[1:]
		}
	}

	return nums
}

func sortArrayByParity(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if nums[j]%2 != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
