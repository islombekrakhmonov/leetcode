package main

import "fmt"

func main() {
	fmt.Println(destCity([][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}))
}

func destCity(paths [][]string) string {

	mapped := make(map[string]int)
	for i :=0; i<len(paths); i++{
		mapped[paths[i][0]]++
	}

	for i :=0; i<len(paths); i++{
		if _, exists := mapped[paths[i][1]]; !exists{
			return paths[i][1]
		}
	}
	
	return ""
}