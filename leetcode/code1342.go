package main

import "fmt"

func main() {
	num := 14
	fmt.Println(numberOfSteps(num))
	
}

func numberOfSteps(num int) int {
	var output int 
	for num >0{
		if num %2 == 0{
			num /=2
			output ++
		} else {
			num -= 1
			output++
		}
	}
	return output
}
