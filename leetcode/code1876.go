package main

import "fmt"

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))
}

func countGoodSubstrings(s string) int {

    var output int 
	for startIndex := 0; startIndex < len(s); startIndex++ {
        for endIndex := startIndex + 1; endIndex <= len(s); endIndex++ {
            substring := s[startIndex:endIndex]
			if len(substring) == 3 {
				if ifAsubstingIsGood(substring) {
					output++
				}
			}
        }
    } 


	return output
}

func ifAsubstingIsGood(s string) bool {
	for i:=0; i<len(s); i++{
		for k:=i+1; k<len(s); k++{
			if s[i] == s[k] {
				return false
			} 
		} 
	}
	return true
}