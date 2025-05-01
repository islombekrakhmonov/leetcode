package main

import "fmt"

func main() {
	// fmt.Println(minOperations("1100"))
	// fmt.Println(minOperations("111000"))
	// fmt.Println(minOperations("1110"))
	// fmt.Println(minOperations("000111"))
	// fmt.Println(minOperations("0000"))
	fmt.Println(minOperations("1111"))
	// fmt.Println(minOperations("1010"))
	// fmt.Println(minOperations("0101"))
}

func minOperations(s string) int {
	count1, count2 := 0, 0

	for i := 0; i < len(s); i++ {
		current := s[i]

		if i%2 == 0 {
			if current != '0' {
				count1++
			}
		} else {
			if current != '1' {
				count1++
			}
		}

		if i%2 == 0 {
			if current != '1' {
				count2++
			}
		} else {
			if current != '0' {
				count2++
			}
		}
	}
	
	if count1 < count2 {
		return count1
	}

	return count2
}
