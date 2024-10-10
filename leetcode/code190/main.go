package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverseBits(43261596))
}

func reverseBits(num uint32) uint32 {

	bits := strconv.FormatUint(uint64(num), 2)

	reversed := string(reverseSlice([]rune(bits)))

	output, _ := strconv.ParseUint(string(reversed), 2, 32)

	return uint32(output)
}

func reverseSlice(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
