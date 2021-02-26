package queue

import (
	"GB/lesson-5/doublelinkedlist"
	"fmt"
)

type Queue struct {
	list *doublelinkedlist.List
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list: &doublelinkedlist.List{},
	}
}
func (s *Queue) Push(elem int) {
	s.list.PushBack(elem)

}
func (s *Queue) Pop() (interface{}, bool) {
	if s.list.Len() == 0 {
		return 0, false
	}
	e := s.list.Front()
	fmt.Print(e.Value)
	s.list.Remove(e)

	return e.Value, true

}
