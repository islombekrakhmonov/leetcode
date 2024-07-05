package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))

}

func mostWordsFound(sentences []string) int {
	var max int

	for _, v := range sentences {
		var output int
		for _, l := range v {
			if string(l) == " " {
				output++
			}
		}
		if output > max {
			max = output + 1
		}
	}
	return max
}
