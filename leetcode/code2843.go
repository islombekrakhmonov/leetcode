package main

import (
	"fmt"
	"time"
)

func main() {
    // Record the start time
    startTime := time.Now()

    // Call the function or run the code you want to measure
    result := countSymmetricIntegers(100, 1000000)

    // Record the end time
    endTime := time.Now()

    // Calculate the execution time
    executionTime := endTime.Sub(startTime)

    // Print the result and execution time
    fmt.Printf("Result: %d\n", result)
    fmt.Printf("Execution Time: %v\n", executionTime)
}

// func countSymmetricIntegers(low int, high int) int {
// 	var output int
//     for i:= low; i<=high; i++{
// 		sum1, sum2 := 0,0
// 		iStr := strconv.Itoa(i)
// 		if  len(iStr) % 2 == 0 {
// 			divided := len(iStr) / 2
// 			for j:=0; j<divided; j++{
// 				val1,_  := strconv.Atoi(string(iStr[j]))
// 				sum1 +=val1
// 			}
// 			for j := len(iStr)-1; j >= divided; j-- {
// 				val2, _ := strconv.Atoi(string(iStr[j]))
// 				sum2 += val2
// 			}		
	
// 			if sum1 == sum2 {
// 				output++
// 			}
// 		}
// 	}
// 	return output
// }

func countSymmetricIntegers(low int, high int) int {
    var output int
    for i := low; i <= high; i++ {
        if isSymmetric(i) {
            output++
        }
    }
    return output
}

func isSymmetric(num int) bool {
    if num < 10 {
        return false // Single-digit numbers are not symmetric
    }

    sum1, sum2 := 0, 0
    digits := []int{} // Store the digits in an array

    // Extract digits from the number and store them in 'digits'
    for num > 0 {
        digit := num % 10
        digits = append(digits, digit)
        num /= 10
    }

    // Calculate the sums of the first and last half of the digits
    n := len(digits)
    for i := 0; i < n/2; i++ {
        sum1 += digits[i]
        sum2 += digits[n-1-i]
    }

    return sum1 == sum2
}