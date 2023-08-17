package main

import "fmt"

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{4,2,3,3,2,5,6}, 3))
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) (count int) {
    for _,v := range hours {
		if v>=target {
			count++
		}
	}

	return 
}