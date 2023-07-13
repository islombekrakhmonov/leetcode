package main

import "fmt"

func main() {
	fmt.Println(theMaximumAchievableX(4,1))
}

func theMaximumAchievableX(num int, t int) int {
	return num + (t*2)
}