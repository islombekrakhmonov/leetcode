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

	fmt.Println(modifiedList([]int{4, 3}, myList1.Next))
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	newList := &ListNode{}
	current := newList
	for head != nil {
		if _, exists := numSet[head.Val]; !exists {
			current.Next = &ListNode{Val: head.Val}
			current = current.Next
		}
		head = head.Next
	}

	return newList.Next
}
