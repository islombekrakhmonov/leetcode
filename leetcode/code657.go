package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("RLUURDDDLU"))
}

func judgeCircle(moves string) bool {

	position := []int{1,0}

	for i:=1;i<len(moves);i++{
		if string(moves[i-1]) == "D" {
			if string(moves[i]) == ""
		}
	}

	if string(moves[0]) == "D" && string(moves[1]) == "D" || string(moves[1]) == "R" || string(moves[1]) == "L"{
		return false
	} 
	if string(moves[0]) == "R" && string(moves[1]) == "D" || string(moves[1]) == "R" || string(moves[1]) == "U" {
		return false
	}
	if string(moves[0]) == "L" && string(moves[1]) == "D" || string(moves[1]) == "U" || string(moves[1]) == "L" {
		return false
	}
	if string(moves[0]) == "U" && string(moves[1]) == "U" || string(moves[1]) == "R" || string(moves[1]) == "L" {
		return false
	}

	return true 
}


