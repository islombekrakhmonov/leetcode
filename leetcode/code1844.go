package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(replaceDigits("a1b2c3d4e"))
}

func replaceDigits(s string) string {
	var output string
	
	for i:=1; i<len(s); i+=2{
		iStr, _ := strconv.Atoi(string(s[i]))
		returned := shift(s[i-1], iStr)
		output += string(s[i-1]) + returned
	}
	if len(s) % 2 !=0 {
		output += string(s[len(s)-1])
	}
	
	
	return output
}

func shift(letter byte, num int) (output string) {
	fmt.Println(string(letter), num)
	var iteration int
	for i := letter; i <= 'z'; i++ {
		output = string(i)
		if iteration == num {
			return output
		}
		iteration ++
	}

	return
}