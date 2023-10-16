package main

import "fmt"

func main() {
	fmt.Println(vowelStrings([]string{"are","amy","u"}, 0,2))
}

func vowelStrings(words []string, left int, right int) int {
	var output int
    
	for i, v := range words {
		lastChar := v[len(v)-1]
		firstChar := v[0]
		if isVowel(firstChar) && isVowel(lastChar) {
			if i >= left && i <= right {
				output++
			}
		}
	}

	return output
}

func isVowel(char byte) bool {
	vowelChar := []byte{'a', 'e', 'i', 'o', 'u'}
    for _, v := range vowelChar {
        if char == v {
            return true
        }
    }
    return false
}



