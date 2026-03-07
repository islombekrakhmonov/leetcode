package main

import "fmt"

func main() {
	fmt.Println(removeDigit("1231", '1'))
}

func removeDigit(number string, digit byte) string {
	var (
		output string
	)

	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			removed := number[:i] + number[i+1:]
			if removed > output {
				output = removed
			}
		}
	}

	return output
}
