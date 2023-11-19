package main

import "fmt"

func main() {
	fmt.Println(checkIfExist([]int{7,1,14,11}))
}

func checkIfExist(arr []int) bool {
	for i := 0; i<len(arr); i++ {
		for j:=0; j<len(arr); j++{
			if arr[i] == arr[j] * 2 && i!= j {
				return true
			}
		}
	}
    
	return false 
}