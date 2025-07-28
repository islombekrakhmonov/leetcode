package main

import "fmt"

func main() {
	fmt.Println(isWinner([]int{5, 10, 3, 2}, []int{6, 5, 7, 3}))
}
func isWinner(player1, player2 []int) int {
	calcScore := func(player []int) int {
		score := 0
		for i := 0; i < len(player); i++ {
			multiplier := 1
			if (i >= 1 && player[i-1] == 10) || (i >= 2 && player[i-2] == 10) {
				multiplier = 2
			}
			score += player[i] * multiplier
		}
		return score
	}

	score1 := calcScore(player1)
	score2 := calcScore(player2)

	if score1 > score2 {
		return 1
	} else if score2 > score1 {
		return 2
	}
	return 0
}
