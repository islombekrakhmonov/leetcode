package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
    var output int
	var length string
	var to []string
	l:= strings.Split(s, " ");
	for _,v := range l{
		if v != ""{
			to = append(to, v)
		}
	}
	fmt.Println(to)
	length = to[len(to)-1];
	output = len(length)
	return output
}