package main

import "fmt"

func main() {
	permutations := Perm([]int{1, 2, 3})
	fmt.Println(permutations)
}

func Perm(a []int) [][]int {
	return perm(a, 0)
}

func perm(a []int, i int) [][]int {
	if i >= len(a) {
		return [][]int{append([]int{}, a...)}
	}

	var result [][]int
	for j := i; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutations := perm(a, i+1)
		result = append(result, permutations...)
		a[i], a[j] = a[j], a[i]
	}

	return result
}