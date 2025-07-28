package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	myList := &ListNode{}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: 1}

	myList.Next = node1
	node1.Next = node2
	node2.Next = node3

	fmt.Println(deleteMiddle(myList))
}

func deleteMiddle(head *ListNode) *ListNode {
	var (
		count        int
		count2       int
		originalHead = head
	)

	for head != nil {
		count++
		head = head.Next
	}

	if count%2 == 0 {
		count = count/2 + 1
	} else {
		count = (count + 1) / 2
	}

	newList := &ListNode{}
	current := newList

	for curr := originalHead; curr != nil; curr = curr.Next {
		count2++
		if count2 == count {
			continue
		}
		node := &ListNode{Val: curr.Val}
		current.Next = node
		current = current.Next
	}

	return newList.Next
}
