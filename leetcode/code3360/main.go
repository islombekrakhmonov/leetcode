package main

import "fmt"

func main() {

	fmt.Println(canAliceWin(10))
}

func canAliceWin(n int) bool {
	if n < 10 {
		return false
	}

	removeCount := 10

	for n >= 0 {
		if n-removeCount < 0 {
			if removeCount%2 == 0 {
				return false
			} else {
				return true
			}
		}
		n -= removeCount
		removeCount--
	}

	return false
}
