package stack

type Stack struct {
	sli []int
}

func NewStack(cap int) *Stack {
	return &Stack{
		sli: make([]int, 0, cap),
	}
}
func (s *Stack) Push(elem int) {
	s.sli = append(s.sli, elem)
}
func (s *Stack) Pop() (int, bool) {
	if len(s.sli) == 0 {
		return 0, false
	}
	elem := s.sli[len(s.sli)-1]
	s.sli = s.sli[:len(s.sli)-1]
	return elem, true
}
