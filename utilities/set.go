package utilities

type Set[K comparable] struct {
	items map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	s := new(Set[K])
	s.items = make(map[K]struct{})
	return s
}

func (s *Set[K]) Add(element K) {
	s.items[element] = struct{}{}
}

func (s *Set[K]) Contains(element K) bool {
	_, ok := s.items[element]
	return ok
}

func (s *Set[K]) Size() int {
	return len(s.items)
}
