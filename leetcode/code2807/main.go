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

	fmt.Println(insertGreatestCommonDivisors(myList1.Next))
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	newList := &ListNode{}
	current := newList

	for curr := head; curr.Next != nil; curr = curr.Next {
		current.Next = &ListNode{Val: curr.Val}
		current = current.Next

		gcd := findGCD(curr.Val, curr.Next.Val)
		current.Next = &ListNode{Val: gcd}
		current = current.Next
	}

	current.Next = &ListNode{Val: head.Val}
	for curr := head; curr.Next != nil; curr = curr.Next {
		current.Next = &ListNode{Val: curr.Next.Val}
	}

	return newList
}

func findGCD(a, b int) int {
	var output, max, min int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	for i := max; i > 0; i-- {
		if min%i == 0 && max%i == 0 {
			output = i
			break
		}
	}
	return output
}
