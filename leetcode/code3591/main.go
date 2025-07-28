package main

import "math"

func main() {
	checkPrimeFrequency([]int{1, 2, 3, 4, 5, 4})
}

func checkPrimeFrequency(nums []int) bool {
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for _, freq := range frequencyMap {
		if isPrime(freq) {
			return true
		}
	}

	return false
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
