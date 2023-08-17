package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
    alphabet := []string{"","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	aToI := make(map[string]string)

	for i,v := range alphabet{
		number := strconv.Itoa(i)
		if i <=9 {
			aToI[number] = v
		} else {
			number = number + "#"
			aToI[number] = v
		}
	}

	str := strings.Split(s, "#")
	fmt.Println(str)


	fmt.Println(aToI)
	return ""
}