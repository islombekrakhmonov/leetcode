package main

import "fmt"

func main() {
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}

func decrypt(code []int, k int) []int {

	l := len(code)
	output := make([]int, l)

	if k == 0 {
		return output
	}

	if k > 0 {
		for i := 0; i < l; i++ {
			sum := 0
			for j := 1; j <= k; j++ {
				nextIndex := (i + j) % l
				sum += code[nextIndex]
			}

			output[i] = sum
		}
	}

	if k < 0 {
		k = -k
		for i := 0; i < l; i++ {
			sum := 0
			for j := 1; j <= k; j++ {
				previousIndex := (i - j + l) % l
				sum += code[previousIndex]
			}

			output[i] = sum
		}
	}

	return output

}
