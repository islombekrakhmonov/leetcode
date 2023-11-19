package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("aA"))
}

func reverseVowels(s string) string {
    output := []byte(s)
	j := len(s)-1
	i := 0

	for i<j {
		for i<j && !isVowel(output[i]){
			i++
		}
		for i<j && !isVowel(output[j]) {
			j--
		}
		output[i], output[j] = output[j], output[i] 
		i++
		j--
	}
	return string(output)
}

func isVowel(char byte) bool {
	vowelChar := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for _, v := range vowelChar {
        if char == v {
            return true
        }
    }
    return false
}

