package main

import "fmt"

func main() {

	myList := &ListNode{}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	myList.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println(reverseList(myList))
}

func reverseList(head *ListNode) *ListNode {

	var slice []int

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	reversed := reverse(slice)

	newList := &ListNode{}
	current := newList

	for _, v := range reversed {
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

func reverse(nums []int) []int {
	var output []int

	for i := len(nums) - 1; i >= 0; i-- {
		output = append(output, nums[i])
	}

	return output
}
