package main

import "fmt"

func main() {
	input := "XI"
	fmt.Println(romanToInt(input))
}

func romanToInt(s string) int {
	var output int

	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	length := len(s)

	if length == 0 {
        return 0
    }

    if length == 1 {
        return characterMap[s[0]]
    }
    output = characterMap[s[length - 1]]
	for i := length - 2; i >= 0; i-- {
        if characterMap[s[i]] < characterMap[s[i+1]] {
            output -= characterMap[s[i]]
        } else {
            output += characterMap[s[i]]
        }
    }

    return output
}

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/