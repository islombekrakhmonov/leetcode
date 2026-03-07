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

	fmt.Println(sortList(myList1.Next))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	newList := &ListNode{}
	current := newList

	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	insertionSort(values)
	for i := 0; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
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
