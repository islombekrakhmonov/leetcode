package main

import "fmt"

func main() {
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
}

func countPoints(rings string) int {
	if len(rings) < 3 {
		return 0
	}

	var count int

	for i:=0; i<=9; i++ {
		blue, red, green := false, false, false
		for j:=0; j<len(rings); j+=2{
			color := rings[j]
			rod := rings[j+1]
			if rod == byte('0'+i) {
				if color == 'B' {
					blue = true
				} else if color == 'R' {
					red = true
				} else if color == 'G' {
					green = true
				}
			}
		}
		if blue && red && green {
			count++
		}
	}

    
	return count
}