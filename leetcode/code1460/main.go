package main

import (
	"strconv"
	"strings"
)

func main() {
	
}


func canBeEqual(target []int, arr []int) bool {
    var count int 
	var notContains int
	
	targetString := strconv.Itoa(target[])
	for _, v := range arr {
		notContains = 0
			if !strings.Contains(target, v){
				notContains ++
			}
		if notContains == 0 {
			count++
		}
	}
	return true
}