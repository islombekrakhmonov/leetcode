package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3,1,2,1}, 5))
}

func processQueries(queries []int, m int) []int {
    var output []int
	var perm []int
	var perm1 []int
	for i:=1; i<=m; i++{
		perm = append(perm, i)
	}

	for _,v := range queries { 
		for p,k := range perm{
			if v == k {
				output = append(output, p)
				perm1 = append([]int{v}, perm[:p]... )
				perm1 = append(perm1, perm[p+1:]...)
				perm = perm1
				break
			}
		}
	} 

	return output
}