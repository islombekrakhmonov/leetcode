package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5,1,5,0,-7}))
}

func largestAltitude(gain []int) int {
    altitude := 0
    max := 0 

	for i:=0; i<len(gain); i++{
        altitude += gain[i]
        if max < altitude {
            max = altitude
        }
    }
    return max
}