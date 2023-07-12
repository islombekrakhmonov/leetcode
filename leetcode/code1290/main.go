package main

import (
	"fmt"
	"strconv"
)


type ListNode struct {
	Val int
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

	fmt.Println(getDecimalValue(myList))
}

func getDecimalValue(head *ListNode) int {
	var bin string
	for head != nil {
		bin += fmt.Sprintf("%v", head.Val)
		head = head.Next
	}

	i, _ := strconv.ParseInt(bin, 2, 64)
	
  return int(i)
}