package stack

type Stack struct {
	items []int
	num   int
}

func (s *Stack) Push(item int) {
	if s.items == nil {
		s.items = make([]int, 0)
	}
	s.items = append(s.items, item)
	s.num++
}

func (s *Stack) Pop() (int, bool) {
	if s.items == nil || s.num == 0 {
		return 0, false
	} else {
		old := s.items[s.num-1]
		s.items = s.items[:s.num-1]
		s.num--
		return old, true
	}
}
