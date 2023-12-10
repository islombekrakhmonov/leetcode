package main

import "fmt"

func main() {
	fmt.Println(kItemsWithMaximumSum(3,3,5,11))
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if numOnes >= k {
		return k
	}

	if numOnes + numZeros >= k {
        return numOnes
    }

    return numOnes - (k-numOnes-numZeros)
}