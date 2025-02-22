package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
}

func judgeCircle(moves string) bool {
	positions := make([]int, 2)

	for _, move := range moves {
		switch move {
		case 'R':
			positions[0]++
		case 'L':
			positions[0]--
		case 'U':
			positions[1]++
		case 'D':
			positions[1]--
		}
	}

	return positions[0] == 0 && positions[1] == 0
}
