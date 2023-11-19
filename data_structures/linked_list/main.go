package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) prepand(n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return 
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head 
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 12}
	node3 := &node{data: 4}
	node4 := &node{data: 2}
	node5 := &node{data: 66}
	node6 := &node{data: 27}
	node7 := &node{data: 45}
	node8 := &node{data: 75}
	myList.prepand(node1)
	myList.prepand(node2)
	myList.prepand(node3)
	myList.prepand(node4)
	myList.prepand(node5)
	myList.prepand(node6)
	myList.prepand(node7)
	myList.prepand(node8)
	
	myList.printListData()
	myList.deleteWithValue(2)
	myList.printListData()
}