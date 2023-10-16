package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1,4,2,5,3}))
}

func sumOddLengthSubarrays(arr []int) int {

	var output int 
	for startIndex := 0; startIndex < len(arr); startIndex++ {
        for endIndex := startIndex + 1; endIndex <= len(arr); endIndex++ {
            subarray := arr[startIndex:endIndex]
			if len(subarray) % 2 != 0 {
				for _,v :=range subarray {
					output += v
				}
			}
        }
    } 


	return output
}