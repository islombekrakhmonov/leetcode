package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(splitWordsBySeparator([]string{"$easy$","$problem$"}, '$'))
}

func splitWordsBySeparator(words []string, separator byte) []string {
	var output []string
	for _,v := range words {
		splitted := strings.Split(v, string(separator))
		for _,k := range splitted {
			if k != ""{
				output = append(output, k)
			}
		}
	}
    
	return output
}