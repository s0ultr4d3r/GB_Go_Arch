package queue

type Queue struct {
	sli []int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		sli: make([]int, 0, cap),
	}
}
func (s *Queue) Push(elem int) {
	s.sli = append(s.sli, elem)
}
func (s *Queue) Pop() (int, bool) {
	if len(s.sli) == 0 {
		return 0, false
	}
	elem := s.sli[0]

	if len(s.sli) == 1 {
		s.sli = nil
		return elem, true
	}

	s.sli = s.sli[1:]
	return elem, true
}
