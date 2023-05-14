package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(47,85))
}
func selfDividingNumbers(left int, right int) []int {
	var numInRange []int
	var output []int
	for x:=left;x<= right;x++{
		numInRange = append(numInRange, x)
	}
	for i:=left;i<=right;i++{
		if isSelfDividing(i){
			output = append(output, i)
		}
		}
	return output
}

func isSelfDividing(i int) bool {
	r := i
	for ; r != 0; r /= 10 {
	  d := r % 10
	  if d == 0 || i%d != 0 {
		return false
	  }
	}
	return true
  }