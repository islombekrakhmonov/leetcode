package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

func maxDepth(s string) int {
	counter, maxCounter := 0,0
    for i:=0; i<len(s); i++{
		if s[i] == '(' {
			counter++
		} else if s[i] == ')' {
			counter--
		}

		if counter > maxCounter {
			maxCounter = counter
		}
	}

	return maxCounter
}