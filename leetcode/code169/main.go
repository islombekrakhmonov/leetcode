package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(isValidName("JohnDoe"))
	fmt.Println(isValidName("Johng'Doe"))
	fmt.Println(isValidName("775675"))
	fmt.Println(isValidName("JohnDoe3"))
	fmt.Println(isValidName("12341"))
	fmt.Println(isValidName("JohnDoe2"))
	fmt.Println(isValidName("JohnDoe-2"))
	fmt.Println(isValidName("12341"))
}

func isValidName(name string) bool {
	// Regular expression to match only letters
	regex := regexp.MustCompile("^[a-zA-ZÀ-ÖØ-öø-ÿ']+([-][a-zA-ZÀ-ÖØ-öø-ÿ']+)*$")
	return regex.MatchString(name)
}
func majorityElement(nums []int) int {
	var candidate, count int

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}

	return candidate
}
