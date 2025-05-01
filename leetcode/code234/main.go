package main

import "fmt"

func main() {
	myList1 := &ListNode{}
	node11 := &ListNode{Val: 1}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 2}
	node14 := &ListNode{Val: 1}
	// node15 := &ListNode{Val: 5}

	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	// node14.Next = node15

	fmt.Println(isPalindrome(myList1.Next))
}

func isPalindrome(head *ListNode) bool {

	var array []int
	for curr := head; curr != nil; curr = curr.Next {
		array = append(array, curr.Val)
	}

	left, right := 0, len(array)-1
	for left < right {
		if array[left] != array[right] {
			return false
		}
		left++
		right--
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
