package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{12,1,12}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    var output []bool
	var demo []int
	var max int
	for i:=0;i< len(candies); i++{
		demo = append(demo, candies[i] + extraCandies)
	}
	for _,v:= range demo{
		max = 0
		for _,l := range candies{
			if v >= l{
				max ++
			} 
		}
		if max == len(candies) {
			output = append(output, true)
		} else {
			output = append(output, false)
		}
	}
	return output
}