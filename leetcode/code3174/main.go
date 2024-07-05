package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(clearDigits("cb34"))
}

func clearDigits(s string) string {
	markingMap := make(map[string]bool)
	var output string

	for i := len(s) - 1; i >= 0; i-- {
		var k = i
		if s[i] >= '0' && s[i] <= '9' {
			iStr := strconv.Itoa(i)
			mark := string(s[i-1]) + iStr
			if s[i-1] > '9' && !markingMap[mark] {
				markingMap[mark] = true
			} else {
				if k > 0 {
					for {
						k--
						kStr := strconv.Itoa(k)
						mark := string(s[k-1]) + kStr
						if k == 0 {
							break
						}
						if s[k-1] > '9' && !markingMap[mark] {
							markingMap[mark] = true
							break
						}
					}
				}
			}
		}
	}

	for i, v := range s {
		iStr := strconv.Itoa(i + 1)
		mark := string(s[i]) + iStr
		if !markingMap[mark] && v > '9' {
			output += string(v)
		}
	}

	return output
}
