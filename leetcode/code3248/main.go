package main

import "fmt"

func main() {

	fmt.Println(finalPositionOfSnake(2, []string{"RIGHT", "DOWN"}))
}

func finalPositionOfSnake(n int, commands []string) int {

	var i, j int

	for _, command := range commands {
		if command == "RIGHT" {
			j++
		} else if command == "DOWN" {
			i++
		} else if command == "UP" {
			i--
		} else if command == "LEFT" {
			j--
		}
	}

	return (i * n) + j
}
