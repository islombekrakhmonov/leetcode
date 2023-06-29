package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("panagram", "napagram"))
}

func isAnagram(s string, t string) bool {
	var count int 
	if len(s) != len(t) {
		return false
	}
    for _,v := range s{
		count = 0
		for _,l := range t{
			if v == l{
				count ++
				t = strings.Replace(t, string(l),"",1)
				fmt.Println(t)
				break
			}
		}
		if count == 0{
			return false
		}
	}
	return true 
}