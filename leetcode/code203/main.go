package main

import "fmt"

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

	fmt.Println(removeElements(myList1.Next, 4))
}

func removeElements(head *ListNode, val int) *ListNode {
	newList := &ListNode{}
	current := newList

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val != val {
			node := &ListNode{Val: curr.Val}
			current.Next = node
			current = current.Next
		}
	}

	return newList.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
