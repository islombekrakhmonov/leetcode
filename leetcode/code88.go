package main

import "fmt"

func main() {
	num1 := []int{0}
	num2 := []int{1}
	merge(num1,0,num2,1)
	fmt.Println(num1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	index :=0
    for i:=m; i<len(nums1); i++{
		if n >0 {
			nums1[i] = nums2[index]
			n--
		}
		index++
	}

	insertionSort(nums1)
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