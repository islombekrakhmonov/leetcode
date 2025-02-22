package main

import "fmt"

func main() {
	fmt.Println(kthCharacter(5))
}

func kthCharacter(k int) byte {

	word := "a"
	for {
		if len(word) >= k {
			break
		}
		generated := ""
		for _, v := range word {
			generated += string(v + 1)
		}
		word += generated
	}

	return byte(word[k-1])
}
