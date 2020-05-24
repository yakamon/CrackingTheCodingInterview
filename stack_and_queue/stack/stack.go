package stack

// Stack is a data structure.
type Stack struct {
	top  *node
	size int
}

type node struct {
	value int
	prev  *node
}

// New returns a new stack.
func New() *Stack {
	return &Stack{nil, 0}
}

// Size returns the size of stack.
func (s *Stack) Size() int {
	return s.size
}

// IsEmpty checks if stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Push adds a new element with the value to the top of stack.
func (s *Stack) Push(value int) {
	n := &node{value, s.top}
	s.top = n
	s.size++
}

// Pop removes the top element from stack s and returns the value.
func (s *Stack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	n := s.top
	s.top = n.prev
	s.size--
	return n.value, true
}

// Peek returns the value of the top element of stack.
func (s *Stack) Peek() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	return s.top.value, true
}
