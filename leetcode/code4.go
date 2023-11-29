package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2,6}, []int{3,4,5}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	insertionSort(nums1)

	l := len(nums1)
	if l % 2 != 0 {
		return float64(nums1[l/2])
	} else {
		sum := float64(nums1[l/2-1] + nums1[l/2])
		return sum/2

	}
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