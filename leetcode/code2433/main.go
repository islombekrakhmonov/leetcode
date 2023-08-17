package main

import "fmt"

func main() {
	fmt.Println(findArray([]int{5,2,0,3,1}))
}

func findArray(pref []int) []int {
    var output []int

	output = append(output, pref[0])
	
	for i:=1;i<len(pref);i++{
		difference := pref[i]^pref[i-1]
		output = append(output, difference)
	}

	return output
}