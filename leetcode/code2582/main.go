package main

import "fmt"

func main() {
	fmt.Println(passThePillow(3, 2))
}

func passThePillow(n int, time int) int {
	cycleLength := 2 * (n - 1)
	t := time % cycleLength

	if t < n {
		return t + 1 
	}

	return (2*n - 1) - t
}
