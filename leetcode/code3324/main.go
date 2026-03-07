package main

import "fmt"

func main() {
	fmt.Println(stringSequence("cab"))
	// output should be ["a", "b", "c", "ca", "caa", "cab"]
}

func stringSequence(target string) []string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	var output []string
	key1, key2 := "a", ""

	diff = (tgt - curr + 26) % 26

	output = append(output, "a")

	for i := 0; i < len(target); i++ {
		for j := 0; j < len(alphabet); j++ {
		}
	}

	return output
}

func contains(s []string, b string) bool {
	for _, value := range s {
		if value == b {
			return true
		}
	}

	return false
}
