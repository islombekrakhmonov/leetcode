package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	numA := new(big.Int)
	numA.SetString(a, 2) // Base 2

	numB := new(big.Int)
	numB.SetString(b, 2) // Base 2

	sum := new(big.Int).Add(numA, numB)

	return sum.Text(2)
}
