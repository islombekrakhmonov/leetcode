package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
}

func divisorGame(n int) bool {

	return n%2 == 0
}
