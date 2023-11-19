package main

import "fmt"

type Stack struct {
	items []byte
}

// Push 
func (s *Stack) Push(i byte) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() {
	l := len(s.items) - 1
	// toRemove := s.items[l]
	s.items = s.items[:l]
}


func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	fmt.Println(myStack)
	myStack.Push(200)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}

