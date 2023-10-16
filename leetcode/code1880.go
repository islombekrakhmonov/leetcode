package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isSumEqual("acb", "cba", "cdb"))
	fmt.Println(isSumEqual1("acb", "cba", "cdb"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	lettersMap := make(map[rune]string)

	letters := "abcdefghij"

	for i,v := range letters {
		iStr := strconv.Itoa(i)
		lettersMap[v] = iStr
	} 

	concatFirst := ""
	concatSecond := ""
	concatTarget := ""


	for _,v := range firstWord {
		concatFirst += lettersMap[v]
	}

	for _,v := range secondWord {
		concatSecond += lettersMap[v]

	}

	for _,v := range targetWord {
		concatTarget += lettersMap[v]
	}

	finalFirst,_ := strconv.Atoi(concatFirst)
	finalSecond,_ := strconv.Atoi(concatSecond)
	finalTarget,_ := strconv.Atoi(concatTarget)
    
	return finalFirst + finalSecond == finalTarget 
}

func isSumEqual1(firstWord string, secondWord string, targetWord string) bool {
    getNumericValue := func(word string) int {
        result := 0
        for _, ch := range word {
            result = result*10 + int(ch-'a')
        }
        return result
    }

    return getNumericValue(firstWord)+getNumericValue(secondWord) == getNumericValue(targetWord)
}
