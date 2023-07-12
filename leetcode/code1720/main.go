package main

import "fmt"

func main() {
	fmt.Println(decode([]int{6,2,7,3}, 4))
}

func decode(encoded []int, first int) []int {
	var output []int 
	output = append(output, first)
	fmt.Println(encoded)

	for i := 0; i < len(encoded); i++ {
		xor := output[i] ^ encoded[i]
		output = append(output, xor)
	}
	return output
}