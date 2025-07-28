package main

import "fmt"

func main() {
	fmt.Println(maxLength([]int{1, 2, 1, 2, 1, 1, 1}))
}

func maxLength(nums []int) int {

	var max int

	for startIndex := 0; startIndex < len(nums); startIndex++ {
		for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			subarray := nums[startIndex:endIndex]
			if len(subarray) > max {
				prod := product(subarray)
				lcm := LCMArray(subarray)
				gcd := findGCD(subarray)
				if prod == (lcm * gcd) {
					max = len(subarray)
				}
			}
		}
	}

	return max
}

func product(nums []int) int {
	var output = 1
	for _, v := range nums {
		output *= v
	}
	return output
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findGCD(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, n := range nums[1:] {
		result = gcd(result, n)
		if result == 1 {
			break
		}
	}
	return result
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func LCMArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, n := range nums[1:] {
		result = lcm(result, n)
	}
	return result
}
