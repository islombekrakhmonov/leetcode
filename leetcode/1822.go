package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{0,0,0,0,66,52,-75,84,-67,-57,68,8,-32,25,20,24,-34,-86,2,44,38,73,17,-22,69,-42,19,84,15,16,-81,-24,-73,-93,9,12,12,29,24,83,57,82,48,-95,57,7,-25,-80,79,33,14,-29,13,-33,18,-35,81}))
}

func arraySign(nums []int) int {
	negative := 0
	for _,v := range nums{
		if v == 0 {
			return 0
		} else if v < 0 {
			negative++
		}
	}
	if negative % 2 != 0 {
		return -1
	} else {
		return 1
	} 
}