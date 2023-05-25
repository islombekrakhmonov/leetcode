package main

import (
	"fmt"
)

func main() {
	fmt.Println(finalValueAfterOperations([]string{"++X","++X","X++"}))
}

func finalValueAfterOperations(operations []string) int {
	var output int
    for i:= 0; i < len(operations); i++{
		if operations[i] == "--X" || operations[i] == "X--"{
			output --
		} 
		if operations[i] == "X++" || operations[i] == "++X"{
			output++
		}
	}
	return output
}