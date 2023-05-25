package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toLowerCase("Hello"))
}

func toLowerCase(s string) string {
    s = strings.ToLower(s)
	return s
}