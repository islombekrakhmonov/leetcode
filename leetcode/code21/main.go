package main

import (
	"fmt"
	"sort"
)

func main() {

	// 1,2,4
	myList1 := &ListNode{}
	node11 := &ListNode{Val: 1}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 4}
	node14 := &ListNode{Val: 4}
	node15 := &ListNode{Val: 5}

	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	node14.Next = node15

	fmt.Println(mergeTwoLists(myList1, myList1))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var array []int

	for curr := list1; curr != nil; curr = curr.Next {
		array = append(array, curr.Val)
	}

	for curr := list2; curr != nil; curr = curr.Next {
		array = append(array, curr.Val)
	}

	sort.Ints(array)

	newList := &ListNode{}
	current := newList

	for _, v := range array {
		node := &ListNode{Val: v}
		current.Next = node
		current = current.Next
	}

	return newList.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
