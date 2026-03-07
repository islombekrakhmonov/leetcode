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
	node14 := &ListNode{Val: 3}

	//18,6,10,3
	myList1.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14

	fmt.Println(insertGreatestCommonDivisors(myList1.Next))
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		gcd := findGCD(curr.Val, curr.Next.Val)
		newNode := &ListNode{Val: gcd}

		newNode.Next = curr.Next
		curr.Next = newNode

		curr = newNode.Next
	}
	return head
}

func findGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return findGCD(b, a%b)
}
