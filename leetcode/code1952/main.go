package main

import "fmt"

func main() {
	fmt.Println(isThree(6))
}

func isThree(n int) bool {
	var output int 
	for i:=1; i<=n; i++{
		if n%i == 0 {
			output++
		}
	}

	return output==3
}