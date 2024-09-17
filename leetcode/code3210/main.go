package main

import "fmt"

func main() {
	fmt.Println(getEncryptedString("dart", 3))
}

func getEncryptedString(s string, k int) string {
	var output string

	for i := range s {
		index := (i + k) % (len(s))
		output += string(s[index])
	}

	return output
}
