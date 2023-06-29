package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1,2,2,1,1,3,3}))
}


func uniqueOccurrences(arr []int) bool {

	occurrence := make(map[int]int)
	
	for _, num := range arr {
        occurrence[num]++
    }

	occurrenceArr := []int{}

	for _, v := range occurrence{
		occurrenceArr = append(occurrenceArr, int(v))
	}

	for i:= 0; i<len(occurrenceArr); i++ {
		for j:= i+1; j<len(occurrenceArr); j++{
			if occurrenceArr[i] == occurrenceArr[j] {
				return false
			} 
		}
	}

	return true 
}