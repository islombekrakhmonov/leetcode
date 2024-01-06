package main

import "fmt"

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}}, 5))
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {

	// records := make(map[int]map[int]bool) // id -> minute -> bool
	// for _, log := range logs {
	// 	id, min := log[0], log[1]
	// 	if records[id] == nil {
	// 		records[id] = make(map[int]bool)
	// 	}
	// 	records[id][min] = true
	// }

	actions := make(map[int]map[int]bool)

	for _, v := range logs {
		userID := v[0]
		action := v[1]
		if actions[userID] == nil {
			actions[userID] = make(map[int]bool)
		}
		fmt.Println(actions)
		actions[userID][action] =true
	}

	output := make([]int, k)
	for _, v := range actions {
		output[len(v)-1]++
	}

	return output
}

func contains(slice []int, element int) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
