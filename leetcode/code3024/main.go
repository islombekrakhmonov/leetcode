package main

import "fmt"

func main() {
	fmt.Println(triangleType([]int{2, 7, 7}))
}

func triangleType(nums []int) string {
	a, b, c := nums[0], nums[1], nums[2]
	if !isTriangle(a, b, c) {
		return "none"
	}

	if a == b && b == c {
		return "equilateral"
	}

	if (a == b && a != c) || (a == c && a != b) || (b == c && b != a) {
		return "isosceles"
	}

	if a != b && b != c && a != c {
		return "scalene"
	}

	return "none"
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
