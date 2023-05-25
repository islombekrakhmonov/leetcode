package main

import "fmt"

func main() {
	acc := [][]int{{1,5},{7,3},{3,5}}
	fmt.Println(maximumWealth(acc))
	
}

func maximumWealth(accounts [][]int) int {
    var max int 
	var sum int 
	for _,v := range accounts{
		sum = 0
		for _,k := range v{
			sum += k
		}
		fmt.Println(sum)
		if sum > max {
			max = sum
		}
	}
	return max
}