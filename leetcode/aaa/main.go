package main

import "fmt"

func getElement(slice []int, index int) int {
    // Convert 1-indexed index to zero-indexed index
    i := index - 1

    // Check if the index is within the valid range
    if i >= 0 && i < len(slice) {
        return slice[i]
    }

    // Handle invalid index (you can choose to return a default value or handle the error differently)
    return 0
}

// Function to set an element in a 1-indexed slice
func setElement(slice []int, index, value int) {
    // Convert 1-indexed index to zero-indexed index
    i := index - 1

    // Check if the index is within the valid range
    if i >= 0 && i < len(slice) {
        slice[i] = value
    }

    // Handle invalid index (you can choose to ignore or handle the error differently)
}

// Example usage:
func main() {
    // Define a 1-indexed slice
    slice := []int{0, 10, 20, 30, 40}

    // Get element at index 3 (1-indexed)
    element := getElement(slice, 3)
    fmt.Println("Element at index 3:", element)

    // Set element at index 2 (1-indexed)
    setElement(slice, 2, 100)

    // Print the updated slice
    fmt.Println("Updated slice:", slice)
}