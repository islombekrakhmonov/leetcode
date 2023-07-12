package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}

func sortSentence(s string) string {
	store := make(map[int]string)
	smth := strings.Split(s, " ")
	var output string
	var numbers []int
	for _,v := range smth {
		convertedInt, _ := strconv.Atoi(string(v[len(v)-1]))
		store[convertedInt] = (string(v[:len(v)-1]))
		numbers = append(numbers, convertedInt)
	}
	
	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		output += store[numbers[i]]
		if i != len(numbers)-1 {
			output += " "
		}
	}

	return output
}