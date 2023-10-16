package main

import "fmt"

func main() {
	fmt.Println(checkString("aababbb"))
}


func checkString(s string) bool {
	var sIndex int 
	var foundB bool
    for sIndex < len(s) {
		if string(s[sIndex]) == "b"{
			foundB = true
		}
		if string(s[sIndex]) == "a" && foundB {
			return false 
		}
		sIndex++
	}
	return true
}