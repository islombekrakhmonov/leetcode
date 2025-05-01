package main

import "fmt"

func main() {
	fmt.Println(reverseDegree("abc"))
}

func reverseDegree(s string) int {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphabetMap := make(map[string]int)

	reversed := reverse((alphabet))

	for i, v := range reversed {
		alphabetMap[v] = i
	}

	var sum int

	for i, v := range s {
		sum += (alphabetMap[string(v)] + 1) * (i + 1)
	}

	return sum
}

func reverse(s []string) []string {
	var output []string
	for i := len(s) - 1; i >= 0; i-- {
		output = append(output, string(s[i]))
	}
	return output
}
