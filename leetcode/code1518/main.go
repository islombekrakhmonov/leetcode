package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(15, 4))
}

func numWaterBottles(numBottles int, numExchange int) int {
	output := numBottles
	empty := numBottles

	for empty >= numExchange {
		full := empty / numExchange
		output += full
		empty = full + (empty % numExchange)
	}

	return output
}
