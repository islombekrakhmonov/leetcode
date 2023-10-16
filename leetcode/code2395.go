package main

import "fmt"

func main() {
	fmt.Println(findSubarrays([]int{1,2,3,4,5}))
}

func findSubarrays(nums []int) bool {
	var sums []int 
	for startIndex := 0; startIndex < len(nums); startIndex++ {
        for endIndex := startIndex + 1; endIndex <= len(nums); endIndex++ {
			sum := 0
            subarray := nums[startIndex:endIndex]
			if len(subarray) == 2 {
				for _,v := range subarray {
					sum += v
				}
				sums = append(sums, sum)
			}
        }
    } 

	for i:=0; i<len(sums); i++{
		for j:= i+1; j<len(sums);j++{
			if sums[i] == sums[j] {
				return true
			}
		}
	}

	return false 
}