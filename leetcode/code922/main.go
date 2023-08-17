package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{3,1,4,6}))
}

func sortArrayByParityII(nums []int) []int {
	l := len(nums)
	oddInd := 1
	evenInd := 0 

	for oddInd < l && evenInd < l{
		for oddInd < l && nums[oddInd]%2 == 1 {
            oddInd += 2
        }

		for evenInd < l && nums[evenInd]%2 == 0 {
            evenInd += 2
        }

		if oddInd < l && evenInd < l {
            nums[oddInd], nums[evenInd] = nums[evenInd], nums[oddInd]
        }
	}
    
	return nums
}