package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood1("leEeetcode"))
}

func makeGood(s string) string {
	found := true
	for found {
		found = false 
		for i:=0; i<len(s)-1; i++ {
			if s[i] == s[i+1]-32 || s[i] == s[i+1]+32 {
				s = s[:i] + s[i+2:]
				found = true
				break
			}
		}
	}
	return s
}

func makeGood1(s string) string {
	ans := make([]rune, 0, len(s))

	for _,c := range s {
    	if len(ans) > 0 && (ans[len(ans)-1]-32 == c || ans[len(ans)-1]+32 == c) {
			ans = ans[0:len(ans)-1]
	  	} else {
			ans = append(ans,rune(c))
	  	}
	}


	return string(ans)
}
