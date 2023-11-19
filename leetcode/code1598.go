package main

import "fmt"

func main() {
	fmt.Println(minOperations([]string{"d1/","d2/","../","d21/","./"}))
}

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() {
	l := len(s.items) - 1
	s.items = s.items[:l]
}

func minOperations(logs []string) int {
	myStack := Stack{}
    for i:=0; i<len(logs); i++ {
		if logs[i] == "./" {
		} else if logs[i] == "../" {
			if len(myStack.items) > 0 {
				myStack.Pop()
			}
		} else {
			myStack.Push(i)
		}
	}

	return len(myStack.items)
}
