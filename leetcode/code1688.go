package main

import "fmt"

func main() {
	fmt.Println(numberOfMatches(14))
}

func numberOfMatches(n int) int {
    var output int
	var total int
	for n>1{
		if n % 2 == 0{
			output = n/2
			n/=2
		} else{
			output = ((n - 1) / 2)
			n = ((n-1)/2)+1
		}
		total += output 
	}
	return total
}