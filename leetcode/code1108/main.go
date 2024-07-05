package main

import "fmt"

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {
	var output string
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			output += replaceDots(address[i])
		} else {
			output += string(address[i])
		}
	}

	return output
}

func replaceDots(byte) string {
	return "[.]"
}
