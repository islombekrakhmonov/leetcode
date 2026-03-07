package main

import "fmt"

func main() {
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	trustMap := make(map[int]int)
	trustedByMap := make(map[int]int)

	for _, v := range trust {
		trustMap[v[1]]++
		trustedByMap[v[0]]++
	}

	for person, count := range trustMap {
		if count == n-1 && trustedByMap[person] == 0 {
			return person
		}
	}

	return -1
}
