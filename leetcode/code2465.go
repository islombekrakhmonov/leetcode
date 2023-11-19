package main

import "fmt"

func main() {
	fmt.Println(distinctAverages([]int{9,5,7,8,7,9,8,2,0,7}))
}

func distinctAverages(nums []int) int {
	insertionSort(nums)

	averages := make(map[float64]int)

	for len(nums) != 0 {
		smallest := nums[0]
		largest  := nums[len(nums)-1]
		nums = nums[1 : len(nums)-1]

		var average float64
		average = float64((smallest+largest))
		average /= 2
		averages[average]++
	}
    
	return len(averages)
}

func insertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}