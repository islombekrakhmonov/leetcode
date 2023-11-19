package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates1(s string) string {
    found := true
	for found {
		found = false 
		for i:=0; i<len(s)-1; i++{
			if s[i] == s[i+1] {
				s = s[:i] + s[i+2:]
				found = true
				break
			}
		}
	}
	return s
}

func removeDuplicates(s string) string {
	ans := make([]rune, 0, len(s))
	for _,c := range s {
	  if 0 < len(ans) && ans[len(ans)-1] == c{
		ans = ans[0:len(ans)-1]
	  } else {
		ans = append(ans,rune(c))
	  }
	}

	return string(ans)
  }

