package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{8,4,6,2,3}))
}

func finalPrices(prices []int) []int {
	var output []int
	for i:=0; i<len(prices); i++{
		for j:=i+1; j<len(prices); j++{
			if prices[i] >= prices[j]{
				newPrice:= prices[i]-prices[j]
				output =append(output, newPrice)
				break
			}
		}
		if len(output) != i+1{
			output= append(output, prices[i])
		}
	}

    return output
}