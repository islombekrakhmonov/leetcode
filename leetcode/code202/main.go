package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isHappy(20))
}


func isHappy(n int) bool {
	seen := make(map[int]bool)
	for {
		sum := 0
		for _,v := range strconv.Itoa(n) {
		num, _ := strconv.Atoi(string(v))
		res := num * num
		sum += res
		}
		if int(sum) == 1 {
			return true
		}
		if seen[sum] {
			return false 
		}

		seen[sum] = true
		n = sum
	}
}