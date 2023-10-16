package main

import "fmt"

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
}

func percentageLetter(s string, letter byte) int {
	var count int
	l := len(s)
	for i := 0; i<l; i++{
		if s[i] == letter{
			count++
		}
	}
	output := count * 100 / l
    
	return output
}