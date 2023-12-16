package main

import (
	"fmt"

)

func main() {
	fmt.Println(isPathCrossing("NESWW"))
}

func isPathCrossing(path string) bool {
	x, y := 0,0
	
	visited := make(map[int]map[int]bool)
	visited[0] = map[int]bool{0: true}

	for i:=0; i<len(path); i++{
		switch path[i] {
		case 'N' : x++
		case 'S' : x--
		case 'E' : y++
		case 'W' : y--
		}

		if _, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
		}

		fmt.Println(visited)

		if visited[x][y] {
			return true
		}

		visited[x][y] = true
	}

	return false
}
