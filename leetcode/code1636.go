package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort([]int{1,1,2,2,2,3}))
}

func frequencySort(nums []int) []int {
	var output []int
	numsMap := make(map[int]int)
    for _,v := range nums {
		numsMap[v]++
	}
	
	var pairs []struct {
        Key   int
        Value int
    }
    for key, value := range numsMap {
        pairs = append(pairs, struct{ Key, Value int }{key, value})
    }

	sort.Slice(pairs, func(i, j int) bool {
        if pairs[i].Value != pairs[j].Value {
            return pairs[i].Value < pairs[j].Value
        }
        // If values are the same, sort by keys in decreasing order
        return pairs[i].Key > pairs[j].Key
    })

	for _, pair := range pairs {
        key, value := pair.Key, pair.Value
        for i := 0; i < value; i++ {
            output = append(output, key)
        }
    }
	
	return output
}