package main

import "fmt"

func main() {
	fmt.Println(checkRecord("ALL"))
}

func checkRecord(s string) bool {

	absenseCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absenseCount++
		}

		if absenseCount > 1 {
			return false
		}

		if s[i] == 'L' && len(s) > i+2 {
			if s[i+1] == 'L' && s[i+2] == 'L' {
				return false
			}
		}
	}
	return true
}
