package main

import "fmt"

func main() {
	fmt.Println(countBeautifulPairs([]int{31,25,72,79,74}))
}

func countBeautifulPairs(nums []int) int {
	var count int 

	for i:=0; i<len(nums)-1; i++{
		for j:=i+1; j<len(nums); j++{
			if gcdEuclidean(nums[i],nums[j]) == 1 {
				count++
			}
		}
	}
    
	return count
}

func gcdEuclidean(a, b int) int {
	if b == 0{
		return a
	}
	return gcdEuclidean(b,a%b)
}
