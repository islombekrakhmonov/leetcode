package main

import (
	"fmt"
	"math/bits"
)

// func main() {
// 	fmt.Println(minChanges(13, 4))
// }

func minChanges(n int, k int) int {

	nBit := fmt.Sprintf("%b", n)
	kBit := fmt.Sprintf("%b", k)

	fmt.Println(nBit, "adsf", kBit)

	for i := 0; i < len(nBit); i++ {
		if kBit[i] == '1' && nBit[i] != '1' {
			return -1
		}
	}

	return 0
}

func equalLengthBits(a, b int) (bitsA, bitsB []int) {
	// Find maximum bit length between the two numbers
	lenA := bits.Len(uint(a))
	lenB := bits.Len(uint(b))
	maxLen := lenA
	if lenB > maxLen {
		maxLen = lenB
	}

	// Convert both numbers to bits with the same length
	bitsA = intToBits(a, maxLen)
	bitsB = intToBits(b, maxLen)
	return
}

// Converts an integer to bits with specified length (pads with leading zeros)
func intToBits(num, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		// Start from MSB (most significant bit)
		mask := 1 << (length - 1 - i)
		if num&mask != 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func main() {
	a := 5  // Binary: 101
	b := 13 // Binary: 1101

	bitsA, bitsB := equalLengthBits(a, b)

	fmt.Printf("%d as bits: %v\n", a, bitsA) // [0 1 0 1]
	fmt.Printf("%d as bits: %v\n", b, bitsB) // [1 1 0 1]

	for i := 0; i < len(bitsA); i++ {
		if bitsA[i] == '1' && bitsB[i] != '1' {
			return -1
		}
	}
}
