package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	var (
		balance int
		output  int
	)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "R" {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			output++
		}
	}

	return output
}

// Loop from left to right maintaining a balance variable when it gets an L increase it by one otherwise decrease it by one.
//Whenever the balance variable reaches zero then we increase the answer by one.
