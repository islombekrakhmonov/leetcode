package main

import "fmt"

func main() {
	fmt.Println(largestInteger([]int{0, 0}, 2))
}

// Hint 1
// Solve the problem for three different cases: k = 1, k = n, and 1 < k < n
// Hint 2
// If k = 1, return the largest element that occurs exactly once in nums
// Hint 3
// If k = n, return the largest element in nums
// Hint 4
// If 1 < k < n, all elements different from nums[0] and nums[n - 1] will occur in more than one subarray of size k. Hence, the answer is the largest of nums[0] and nums[n - 1] if they both occur exactly once in the array. If one of them occurs more than once, return the other. If both of them occur more than once, return -1.

func largestInteger(nums []int, k int) int {

	var output = -1
	var countMap = make(map[int]int)

	if k == 1 {
		for _, v := range nums {
			countMap[v]++
		}

		for k, v := range countMap {
			if v == 1 && k > output {
				output = k
			}
		}

		return output
	} else if k == len(nums) {
		for _, v := range nums {
			if v > output {
				output = v
			}
		}
		return output
	} else {
		num0, numLast := false, false
		if nums[0] == nums[len(nums)-1] {
			return output
		}

		for i := 1; i < len(nums)-1; i++ {
			if nums[i] == nums[0] {
				num0 = true
			}
			if nums[i] == nums[len(nums)-1] {
				numLast = true
			}
		}

		if num0 && numLast {
			output = -1
		} else if !num0 && !numLast {
			if nums[0] > nums[len(nums)-1] {
				output = nums[0]
			} else {
				output = nums[len(nums)-1]
			}
		} else if num0 {
			output = nums[len(nums)-1]
		} else if numLast {
			output = nums[0]
		}
	}

	return output
}
