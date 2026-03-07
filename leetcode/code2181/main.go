package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// [0,3,1,0,4,5,2,0]
func main() {
	myList1 := &ListNode{}
	node11 := &ListNode{Val: 0}
	node12 := &ListNode{Val: 3}
	node13 := &ListNode{Val: 1}
	node14 := &ListNode{Val: 0}
	node15 := &ListNode{Val: 4}
	node16 := &ListNode{Val: 5}
	node17 := &ListNode{Val: 2}
	node18 := &ListNode{Val: 0}

	//18,6,10,3
	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	node14.Next = node15
	node15.Next = node16
	node16.Next = node17
	node17.Next = node18

	fmt.Println(mergeNodes(myList1.Next))
}

func mergeNodes(head *ListNode) *ListNode {

	newList := &ListNode{}
	current := newList
	sum := 0
	head = head.Next

	for head != nil {
		if head.Val != 0 {
			sum += head.Val
		} else {
			current.Next = &ListNode{Val: sum}
			current = current.Next
			sum = 0
		}
		head = head.Next
	}

	return newList.Next
}
