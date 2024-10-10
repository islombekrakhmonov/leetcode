package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digitSum("11111222223", 3))
}

func digitSum(s string, k int) string {

	for len(s) > k {
		var groups []string
		for i := 0; i < len(s); i += k {
			end := i + k
			if end > len(s) {
				end = len(s)
			}
			groups = append(groups, s[i:end])
		}

		s = ""
		for _, group := range groups {
			var sum int
			for _, numStr := range group {
				num, _ := strconv.Atoi(string(numStr))
				sum += num
			}
			sumStr := strconv.Itoa(sum)
			s += sumStr
		}
	}

	return s
}
