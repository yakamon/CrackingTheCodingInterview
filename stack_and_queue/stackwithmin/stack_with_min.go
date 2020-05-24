package stackwithmin

// Stack is stack data structure that holds the min value
type Stack struct {
	top    *node
	length int
}

type node struct {
	value    int
	minValue int
	prev     *node
}

// New initializes a new Stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Len returns the length of stack s
func (s *Stack) Len() int {
	return s.length
}

// Min returns the minimum value of stack s
func (s *Stack) Min() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	return s.top.minValue, true
}

// Push adds a new element to the top of stack s
func (s *Stack) Push(value int) {
	min := value
	if s.top != nil && min > s.top.minValue {
		min = s.top.minValue
	}
	n := &node{value, min, s.top}
	s.top = n
	s.length++
}

// Pop pops an element off the top of stack s
func (s *Stack) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value, true
}

// Peek gets the value of the top element of stack s
func (s *Stack) Peek() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	return s.top.value, true
}
