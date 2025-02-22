package main

import (
	"fmt"
	"sort"
	"strconv"
)

// {3, 2}, {4, 1}
func main() {
	fmt.Println(mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}}))

	tgChatId := -1002272651871

	fmt.Println(tgChatId)

	tgChatIdStr := fmt.Sprintf("%d", tgChatId)

	fmt.Println(tgChatIdStr)

	tgChatIdInt, _ := strconv.Atoi(tgChatIdStr)
	fmt.Println(tgChatIdInt)

	fmt.Println(tgChatIdInt == tgChatId)

	chatIdInt, _ := strconv.ParseInt(tgChatIdStr, 10, 64)

	fmt.Println(chatIdInt)

}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {

	idSumMap := make(map[int]int)

	for _, num := range nums1 {
		idSumMap[num[0]] += num[1]
	}
	for _, num := range nums2 {
		idSumMap[num[0]] += num[1]
	}

	result := make([][]int, 0, len(idSumMap))
	for id, sum := range idSumMap {
		result = append(result, []int{id, sum})
	}
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })

	return result
}
