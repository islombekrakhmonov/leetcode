package main

import "fmt"

func main() {
	fmt.Println(kthCharacter(20))
}

func kthCharacter(k int) byte {
	return findChar(k)
}

func findChar(k int) byte {
	if k == 1 {
		return 'a'
	}

	// Find the length of the block that contains k
	length := 1
	for length < k {
		length *= 2
	}

	half := length / 2

	fmt.Println("length", length)
	fmt.Println("half", half)

	if k <= half {
		return findChar(k) // still in the first half
	} else {
		c := findChar(k - half) // map back to first half
		return c + 1            // shifted by 1
	}
}
