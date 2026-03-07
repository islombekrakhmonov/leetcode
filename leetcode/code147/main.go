package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	myList1 := &ListNode{}
	node11 := &ListNode{Val: 4}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 1}
	node14 := &ListNode{Val: 3}
	// node15 := &ListNode{Val: 4}
	// node16 := &ListNode{Val: 5}
	// node17 := &ListNode{Val: 2}
	// node18 := &ListNode{Val: 0}

	//18,6,10,3
	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	// node14.Next = node15
	// node15.Next = node16
	// node16.Next = node17
	// node17.Next = node18

	fmt.Println(modifiedList(myList1.Next))

	num1 := 3
	num2 := 4
	num3 := 5

	num := (num1 * 100) + (num2 * 10) + num3
	fmt.Println(num)

	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num /= 10
	}
}

func modifiedList(head *ListNode) *ListNode {
	newList := &ListNode{}
	current := newList

	var array []int

	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	insertionSort(array)

	for i := 0; i < len(array); i++ {
		current.Next = &ListNode{Val: array[i]}
		current = current.Next
	}

	return newList.Next
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
