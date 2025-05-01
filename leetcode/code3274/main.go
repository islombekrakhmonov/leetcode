package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(checkTwoChessboards("a1", "c3"))
}

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {

	coor1FirstChar := int(coordinate1[0])
	coor1SecondChar, _ := strconv.Atoi(string(coordinate1[1]))
	coor2SecondChar, _ := strconv.Atoi(string(coordinate2[1]))
	coor2FirstChar := int(coordinate2[0])

	return (coor1FirstChar+coor1SecondChar)%2 == (coor2SecondChar+coor2FirstChar)%2
}
