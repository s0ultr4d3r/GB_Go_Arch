package queue

import "GB/lesson-5/linkedlist"

type Queue struct {
	list *linkedlist.List
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list: &linkedlist.List{},
	}
}
func (s *Queue) Push(elem int) {
	node := &linkedlist.Node{
		Data: elem,
	}
	s.list.Append(node)
}
func (s *Queue) Pop() (int, bool) {
	if s.list.Len() == 0 {
		return 0, false
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head())
	return elem, true

}
