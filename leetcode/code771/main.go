package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("A","aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	var output int
    for i := 0; i<len(jewels); i++{
		for k:= 0; k<len(stones); k++{
			if jewels[i] == stones[k]{
				output ++
			}
		}
	}
	return output
}