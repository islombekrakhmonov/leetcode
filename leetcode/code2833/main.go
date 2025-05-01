package main

import "fmt"

func main() {
	fmt.Println(furthestDistanceFromOrigin("L_RL__R"))
}

func furthestDistanceFromOrigin(moves string) int {
	undescore, l, r := 0, 0, 0

	for _, move := range moves {
		switch move {
		case '_':
			undescore++
		case 'L':
			l++
		case 'R':
			r++
		}
	}

	if l > r {
		return l - r + undescore
	}

	return r - l + undescore
}
