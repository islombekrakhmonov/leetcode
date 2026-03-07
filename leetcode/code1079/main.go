package main

import "fmt"

func main() {

	fmt.Println(numTilePossibilities("AAB"))
}

func numTilePossibilities(tiles string) int {
	tilesCount := make(map[string]int)

	perm := permute([]byte(tiles))

	for _, char := range tiles {
		tilesCount[string(char)]++
	}

	for _, v := range perm {
		tilesCount[string(v)]++
	}

	fmt.Println(tilesCount)

	return len(tilesCount)
}

func permute(nums []byte) [][]byte {
	var result [][]byte
	var backtrack func(int)

	backtrack = func(first int) {
		// if all positions fixed, add a copy to results
		if first == len(nums) {
			perm := make([]byte, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := first; i < len(nums); i++ {
			// swap current element with the first
			nums[first], nums[i] = nums[i], nums[first]
			// recursively fix the rest
			backtrack(first + 1)
			// backtrack (undo swap)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)
	return result
}

func backtrack(letterCount map[rune]int) int {
	count := 0

	for ch, cnt := range letterCount {
		if cnt > 0 {
			// Use this letter
			letterCount[ch]--
			count += 1 + backtrack(letterCount) // count this sequence and all following sequences

			// Backtrack
			letterCount[ch]++
		}
	}

	return count
}
