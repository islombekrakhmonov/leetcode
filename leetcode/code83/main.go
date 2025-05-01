package main

import (
	"fmt"
	"sort"
)

func main() {
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

	fmt.Println(deleteDuplicates(myList1))
}

func deleteDuplicates(head *ListNode) *ListNode {
	valueMap := make(map[int]bool)

	for curr := head; curr != nil; curr = curr.Next {
		valueMap[curr.Val] = true
	}

	var array []int
	for value := range valueMap {
		array = append(array, value)
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
