package main

import (
	"fmt"
	"strconv"
)

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

	fmt.Println(doubleIt(myList1.Next))

	num1 := 3
	num2 := 4
	num3 := 5

	num := (num1 * 100) + (num2 * 10) + num3
	fmt.Println(num)

	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num /= 10
	}
}

func doubleIt(head *ListNode) *ListNode {
	newList := head
	current := newList

	var numStr string

	for head != nil {
		numStr += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}
	num, _ := strconv.Atoi(numStr)
	if num == 0 {
		return &ListNode{Val: 0}
	}
	num *= 2

	numArray := []int{}

	for num >= 0 {
		digit := num % 10
		numArray = append(numArray, digit)
		num /= 10
	}

	numArray = reverse(numArray)

	for _, num := range numArray {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return newList.Next
}

func reverse(nums []int) []int {
	var output []int

	for i := len(nums) - 1; i >= 0; i-- {
		output = append(output, nums[i])
	}

	return output
}
