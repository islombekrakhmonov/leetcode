package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(s string, c byte) []int {
	l := len(s)
    output := make([]int, l)

	prev := -l 
	next := 2 * l

    for i := 0; i < l; i++ {
        if s[i] == c {
            prev = i
        }
        output[i] = i - prev
    }

    for i := l - 1; i >= 0; i-- {
        if s[i] == c {
            next = i
        }
        if output[i] > next - i {
			output[i] = next-i
		} 
    }
	return output
}

