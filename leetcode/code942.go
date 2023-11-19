package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
}

func diStringMatch(s string) []int {
	var output []int

	l := len(s)
	fmt.Println(l)
	ifI := 0

	for i:=0; i<=l; i++{
		if string(s[i]) == "D" {
			output = append(output, l)
			l--
		} else {
			output = append(output, ifI)
			ifI++
		}
		fmt.Println(i)
	}

	// if s[len(s)-1] == 'D' {
	// 	output = append(output, l)
	// } else {
	// 	output = append(output, ifI)
	// }


	return output
}