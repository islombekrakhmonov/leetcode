package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(RandInt(-10, 10))
}

func sumZero(n int) []int {
    var output []int
	for i := 0; i < n; i++ {
		output = append(output, rand.Intn(10))
	}
	for i := 0; i < len(output); i++ {

	}
	return output
}

func RandInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}