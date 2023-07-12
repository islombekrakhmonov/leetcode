package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17,18,5,4,6,1}))
}

func replaceElements(arr []int) []int {
	for i:=0; i<len(arr); i++ {
		var max = 0
		for j:= i+1; j<len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				arr[i] = max
			}

		} 
	}

	arr[len(arr)-1]  = -1
	return arr
}
