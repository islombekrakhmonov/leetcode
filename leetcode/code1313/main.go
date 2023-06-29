package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1,2,3,4}))
}

func decompressRLElist(nums []int) []int {
	var output []int
    for i:=0; i<len(nums); i++{
		if i % 2 == 0 {
			for k:=0; k<nums[i]; k++{
				output = append(output, nums[i+1])
			}
		}
	}
	return output
}

//You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
func checkStraightLine(coordinates [][]int) bool {
	var x1, y1, x2, y2 int
	for i:=0; i<len(coordinates)-1; i++{
		x1 = coordinates[i][0]
		y1 = coordinates[i][1]
		x2 = coordinates[i+1][0]
		y2 = coordinates[i+1][1]
		if (x2-x1)*(y2-y1) != 0 {
			return false
		}
	}
	return true
}

