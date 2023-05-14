package main

import (
	"fmt"
)

func main() {
	s:= "codeleet"
	ins := []int{4,5,6,7,0,2,1,3}
	fmt.Println(restoreString(s,ins))
}

func restoreString(s string, indices []int) string {
	output := make([]byte, len(s))
	for i := 0; i < len(s); i++{
		output[indices[i]] = s[i]
	}
	return string(output)
}