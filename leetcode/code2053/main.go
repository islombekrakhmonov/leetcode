package main

import (
	"fmt"
)

func main() {
	fmt.Println(kthDistinct([]string{"d","b","c","b","c","a"},2))
}

func kthDistinct(arr []string, k int) string {
    mapped := make(map[string]int)
	distinctOrder := []string{}

	for _,v := range arr{
		mapped[v]++
	}

	for _, v := range arr {
        if mapped[v] == 1 {
            distinctOrder = append(distinctOrder, v)
        }
    }

	if k <= len(distinctOrder) {
        return distinctOrder[k-1] 
    }

	return ""
}