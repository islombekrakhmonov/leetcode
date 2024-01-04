package main

import "fmt"

func main() {
	fmt.Println(maxSum([]int{31,25,72,79,74}))
	fmt.Println(hasSameHighestDigits(43, 56))
}

func maxSum(nums []int) int {
	output := -1

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if hasSameHighestDigits(nums[i], nums[j]) {
				if nums[i]+nums[j] > output {
					output = nums[i] + nums[j]
				}
			}
		}
	}

	return output
}

func getDigits(number int) []int {
	if number == 0 {
		return []int{0}
	}

	var digits []int
	for number > 0 {
		digit := number % 10
		digits = append(digits, digit)
		number /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func hasSameHighestDigits(num1, num2 int) bool {
	
	digits := func(number int) int{
		max := 0
		for number > 0 {
			digit := number % 10
			if max < digit  {
				max = digit
			}
			number /= 10
		}
		return max
	}

	num1Max := digits(num1)
	num2Max := digits(num2)
	fmt.Println(num1Max, num2Max)
	
	return num1Max == num2Max
}
