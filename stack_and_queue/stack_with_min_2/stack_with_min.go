package stackwithmin2

// StackWithMin is stack data structure that holds min value
type StackWithMin struct {
	minStack  *Stack
	mainStack *Stack
}

// NewStackWithMin initializes a new StackWithMin instance
func NewStackWithMin() *StackWithMin {
	return &StackWithMin{NewStack(), NewStack()}
}

// Size returns the size of stack
func (s *StackWithMin) Size() int {
	return s.mainStack.Size()
}

// Push adds an element
func (s *StackWithMin) Push(value int) {
	if min, ok := s.minStack.Peek(); ok && value < min {
		s.minStack.Push(value)
	}
	s.mainStack.Push(value)
}

// Pop pops an element
func (s *StackWithMin) Pop() (int, bool) {
	if s.mainStack.IsEmpty() {
		return 0, false
	}
	value, _ := s.mainStack.Pop()
	if min, _ := s.minStack.Peek(); value == min {
		s.minStack.Pop()
	}
	return value, true
}

// Peek returns the value of the top element
func (s *StackWithMin) Peek() (int, bool) {
	return s.mainStack.Peek()
}

// Min returns the min value
func (s *StackWithMin) Min() (int, bool) {
	return s.minStack.Peek()
}

// Stack is stack data structure
type Stack struct {
	top  *node
	size int
}

type node struct {
	value int
	prev  *node
}

// NewStack returns a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Size returns the size of stack s
func (s *Stack) Size() int {
	return s.size
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Push adds a new element with the value to the top of stack s
func (s *Stack) Push(value int) {
	n := &node{value, s.top}
	s.top = n
	s.size++
}

// Pop removes the top element from stack s and returns the value
func (s *Stack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	n := s.top
	s.top = n.prev
	s.size--
	return n.value, true
}

// Peek returns the value of the top element of stack s
func (s *Stack) Peek() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	return s.top.value, true
}
