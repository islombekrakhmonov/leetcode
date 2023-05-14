package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
	
}

func fizzBuzz(n int) []string {
	var output []string
	fizz := "Fizz"
	buzz := "Buzz"
	fizzBuzz := "FizzBuzz"

	for i := 1; i <= n; i++ {
		if i % 3 != 0 && i % 5 != 0 {
			output = append(output, strconv.Itoa(i))
		}  else if i % 3 == 0 && i % 5 == 0 {
			output = append(output, fizzBuzz)
		}  else if i % 3 == 0 {
			output = append(output, fizz)
		}  else if i % 5 == 0 {
			output = append(output, buzz)
		} 
	} 
	return output
}

// if i%3 == 0 {
// 	output = append(output, fizz)
// } else {
// 	output = append(output, buzz)
// }