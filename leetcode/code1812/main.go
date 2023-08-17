package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(squareIsWhite("a1"))

}

func squareIsWhite(coordinates string) bool {
    allCoordinates := make(map[string]bool)
	letters := []string{"a","b","c","d","e","f","g","h"}

	for i:=0;i<len(letters);i++{
		for j:=1;j<=len(letters);j++{
			jStr := strconv.Itoa(j)
			if (i+j)%2 == 0 {
				allCoordinates[letters[i]+jStr] = false
			} else {
				allCoordinates[letters[i]+jStr] = true
			}
		}
	}
	fmt.Println(allCoordinates)

	return allCoordinates[coordinates]
}