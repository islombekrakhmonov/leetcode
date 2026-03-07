package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	fives := 0
	tens := 0
	twenties := 0

	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			tens++
			fives--
		case 20:
			twenties++
			if tens > 0 {
				tens--
				fives--
			} else {
				fives -= 3
			}
		}
		if fives < 0 || tens < 0 || twenties < 0 {
			return false
		}
	}

	return true
}
