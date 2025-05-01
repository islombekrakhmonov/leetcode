package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	myList1 := &ListNode{}
	node11 := &ListNode{Val: 18}
	node12 := &ListNode{Val: 6}
	node13 := &ListNode{Val: 10}

	//18,6,10,3
	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13

	fmt.Println(deleteNode(*myList1))

}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	} else {
		node = nil
	}

	return
}
