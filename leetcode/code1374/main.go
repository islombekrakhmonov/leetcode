package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateTheString(17))
}

func generateTheString(n int) string {
    s := strings.Repeat("a", n-1)
    if n%2 == 0 {
        s += "b"
    } else {
        s += "a"
    }
    return s
}