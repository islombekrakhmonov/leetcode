package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(s string) []int {
	var output []int

	lastIndex := make(map[byte]int)

	for i:=0; i<len(s);i++{
		lastIndex[s[i]] = i
	}
	
	start, end := 0,0
	for i:=0; i<len(s); i++{

		if lastIndex[s[i]] > end {
            end = lastIndex[s[i]]
        }
		if i == end {
            output = append(output, end-start+1)
            start = i + 1
        }
	}

	return output
}
