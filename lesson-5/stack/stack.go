package stack

import (
	"GB/lesson-5/doublelinkedlist"
	"fmt"
)

type Stack struct {
	sli *doublelinkedlist.List
}

func NewStack(cap int) *Stack {
	return &Stack{
		sli: doublelinkedlist.New(),
	}
}
func (s *Stack) Push(elem int) {
	s.sli.PushBack(elem)
}
func (s *Stack) Pop() (interface{}, bool) {
	if s.sli.Len() == 0 {
		return 0, false
	}
	e := s.sli.Back()
	fmt.Print(e.Value)
	s.sli.Remove(e)
	return e.Value, true
}
