package main

import "fmt"

func main() {
	nums:= []int{1,2,4}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
    var ans []int
	ans = nums
	for _,v := range ans{
		ans = append(ans, v)
	}
	return ans
}