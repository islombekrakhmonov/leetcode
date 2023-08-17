package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getNoZeroIntegers(1010))
}


func getNoZeroIntegers(n int) (output []int) {
    for i:=1;i<n;i++{
		for j:=1; j<n;j++{
			if i+j == n && i % 10 != 0 && j % 10 != 0 {
				iStr:= strconv.Itoa(i)
				jStr := strconv.Itoa(j)
				if !strings.Contains(jStr, "0") && !strings.Contains(iStr, "0") {
					output = append(output, i)
					output = append(output, j)
					return
				}
			}
		}
	}

	return
}