package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

func main() {
	count := 100000000
	resultChan := make(chan string)
	done := make(chan struct{})
	generated := make(map[string]int)

	// Goroutine to generate random strings concurrently
	go func() {
		defer close(done)
		for i := 0; i < count; i++ {
			generatedString := GenerateOrderID()
			resultChan <- generatedString
		}
	}()

	// Goroutine to check for duplicates
	go func() {
		for generatedString := range resultChan {
			generated[generatedString]++
		}
	}()

	// Wait for all goroutines to finish
	<-done

	// Print duplicates and their counts
	for key, value := range generated {
		if value > 1 {
			fmt.Printf("Duplicate found: %s, Total duplicates: %d\n", key, value)
		}
	}
}

const (
	Lower = iota + 1
	Upper
	Number
	UpperNumber
	LowerUpper
	LowerNumber
	LowerUpperNumber
)

func GenerateOrderID() string {
	id := uuid.New()
	// Convert UUID to string and remove hyphens
	idString := strings.ReplaceAll(id.String(), "-", "")
	// Take the first 8 characters of the UUID
	return idString[:8]
}

func GenerateRandomString(length int, cmd int) string {
	var letterBytes string

	switch cmd {
	case Lower:
		letterBytes = "abcdefghijklmnopqrstuvwxyz"
	case Upper:
		letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case Number:
		letterBytes = "0123456789"
	case LowerUpperNumber:
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case UpperNumber:
		letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case LowerNumber:
		letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	case LowerUpper:
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	default:
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func diagonalSum(mat [][]int) int {
	n, res := len(mat), 0

	for i := 0; i < n; i++ {
		res += mat[i][i]

		res += mat[i][n-i-1]
	}

	if n%2 != 0 {
		res -= mat[n/2][n/2]
	}

	return res
}
