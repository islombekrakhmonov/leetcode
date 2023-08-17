package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"-60","D","-36","30","13","C","C","-33","53","79"}))
}

func calPoints(operations []string) (output int) {
	stack := make([]int, 0)
	for i:=0; i<len(operations); i++{

		val, ok := strconv.Atoi(operations[i]) 
		if ok == nil{
			stack = append(stack, val)
		} 

		if string(operations[i]) == "C" {
			stack = stack[:len(stack)-1]
		} 
		if string(operations[i]) == "D" {
			double := 2 * stack[len(stack)-1]
			stack = append(stack, double)
		}
		if string(operations[i]) == "+" {
			sum := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, sum)
		}
		fmt.Println(stack)
	}

	for _,v := range stack{
		output += v
	}

	return 
}