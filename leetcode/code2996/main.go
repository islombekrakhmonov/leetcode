package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(missingInteger([]int{4, 5, 6, 7, 8, 8, 9, 4, 3, 2, 7}))
	fmt.Println(IsValidName("John*Doe"))
	fmt.Println(IsValidName("John#Doe"))
	fmt.Println(IsValidName("John%Doe-"))
	fmt.Println(IsValidName("John^"))
	fmt.Println(IsValidName("Johnd'Oe"))
	fmt.Println(IsValidName("JohnDoe1"))
	fmt.Println(IsValidName("John Doe2"))
	fmt.Println(IsValidName("John Doe "))

}

// To find the longest sequential prefix, iterate from left to right. For a fixed i, if nums[i] != nums[i - 1] + 1 then the longest sequential prefix ends at i - 1.

func missingInteger(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] != nums[i-1]+1 {
			break
		}
		sum += nums[i]
	}
	for {
		if !contains(nums, sum) {
			return sum
		} else {
			sum++
		}
	}
}

func contains(slice []int, x int) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}

	return false
}

func IsValidName(name string) bool {
	if strings.TrimSpace(name) == "" {
		return false
	}

	if len(name) < 2 || len(name) > 35 {
		return false
	}
	// Regular expression to match letters, "-" and empty spaces
	regex := regexp.MustCompile("^[a-zA-Z\\s-]*[a-zA-Z][a-zA-Z\\s-]*$")
	return regex.MatchString(name)
}
