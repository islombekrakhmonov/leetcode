package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortPeople([]string{"Mary","John","Emma"}, []int{180,165,170}))
}

func sortPeople(names []string, heights []int) []string {
	names1 := make(map[int]string)
	names2 := make(map[string]int)
	var output []string

    for i:=0; i<len(names); i++{
		names1[heights[i]] = names[i]  
		names2[names[i]] = heights[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	for _,v := range heights{
		if val, ok := names1[v]; ok {
			output = append(output, val)
		}
	}
	return output
}