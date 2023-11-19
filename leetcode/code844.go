package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
}



func backspaceCompare(s string, t string) bool {

	
	sResult, tResult := "", ""

    for i:=0; i<len(s); i++{
		if s[i] != '#' {
			sResult+= string(s[i])
		} else if len(sResult) > 0 {
			sResult = sResult[:len(sResult)-1]
		}
	}

	for i:=0; i<len(t); i++{
		if t[i] != '#' {
			tResult+= string(t[i])
		} else if len(tResult) > 0 {
			tResult = tResult[:len(tResult)-1]
		}
	}

	return sResult == tResult
}

type Stack struct {
	items []byte
}

func (s *Stack) Push(i byte) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() {
	l := len(s.items) - 1
	// toRemove := s.items[l]
	s.items = s.items[:l]
}
