package stackset

// StackSet is set of stacks
type StackSet struct {
	stacks []*Stack
	percap int
	length int
}

// NewStackSet returns a new StackSet instance
func NewStackSet(percap int) *StackSet {
	var stacks []*Stack
	return &StackSet{stacks, percap, 0}
}

// Len returns total number of elements
func (s *StackSet) Len() int {
	return s.length
}

// NumOfStacks returns number of stacks
func (s *StackSet) NumOfStacks() int {
	return len(s.stacks)
}

// Push adds a new element
func (s *StackSet) Push(value int) {
	if last := s.StackOfIndex(len(s.stacks) - 1); last != nil && !last.IsFull() {
		last.Push(value)
		s.stacks[len(s.stacks)-1] = last
	} else {
		stack := NewStack(s.percap)
		stack.Push(value)
		s.stacks = append(s.stacks, stack)
	}
	s.length++
}

// Pop pops last element of last stack
func (s *StackSet) Pop() (int, bool) {
	last := s.StackOfIndex(len(s.stacks) - 1)
	if last == nil {
		return 0, false
	}

	lastValue := last.Pop()
	if last.Size() == 0 {
		s.removeStack(len(s.stacks) - 1)
	}
	s.length--

	return lastValue, true
}

// PopAt pops value from any stack in the stack-set
func (s *StackSet) PopAt(stackIndex int) (int, bool) {
	return s.leftShift(stackIndex, true)
}

func (s *StackSet) leftShift(stackIndex int, removeTop bool) (int, bool) {
	stack := s.StackOfIndex(stackIndex)
	if stack == nil {
		return 0, false
	}

	var removedItem int
	if removeTop {
		removedItem = stack.Pop()
	} else {
		removedItem = stack.RemoveBottom()
	}

	if stack.IsEmpty() {
		s.removeStack(stackIndex)
	} else if s.NumOfStacks() > stackIndex+1 {
		v, _ := s.leftShift(stackIndex+1, false)
		stack.Push(v)
		s.stacks[stackIndex] = stack
	}

	return removedItem, true
}

// Peek returns last element of last stack
func (s *StackSet) Peek() (int, bool) {
	last := s.StackOfIndex(len(s.stacks) - 1)
	if last == nil {
		return 0, false
	}
	return last.Peek(), true
}

// StackOfIndex gets stack by index of stacks
func (s *StackSet) StackOfIndex(stackIndex int) *Stack {
	if stackIndex < 0 || len(s.stacks)-1 < stackIndex {
		return nil
	}
	return s.stacks[stackIndex]
}

func (s *StackSet) removeStack(stackIndex int) bool {
	if stackIndex >= s.NumOfStacks() {
		return false
	}

	if stackIndex == s.NumOfStacks()-1 {
		s.stacks = s.stacks[:stackIndex]
	} else {
		s.stacks = append(s.stacks[:stackIndex], s.stacks[stackIndex+1:]...)
	}
	return true
}

// Stack is stack data structure
type Stack struct {
	top    *node
	bottom *node
	size   int
	cap    int
}

type node struct {
	value int
	above *node
	below *node
}

func newNode(value int) *node {
	return &node{value, nil, nil}
}

// NewStack returns a new stack
func NewStack(cap int) *Stack {
	return &Stack{nil, nil, 0, cap}
}

// Size returns the size of stack s
func (s *Stack) Size() int {
	return s.size
}

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// IsFull checks if stack is full
func (s *Stack) IsFull() bool {
	return s.size == s.cap
}

// Push adds a new element with the value to the top of stack s
func (s *Stack) Push(value int) bool {
	if s.size >= s.cap {
		return false
	}

	s.size++
	n := newNode(value)
	if s.size == 1 {
		s.bottom = n
	}
	if s.top != nil {
		s.top.above = n
	}
	if n != nil {
		n.below = s.top
	}
	s.top = n

	return true
}

// Pop removes the top element from stack s and returns the value
func (s *Stack) Pop() int {
	t := s.top
	s.top = t.below
	s.size--
	return t.value
}

// Peek returns the value of the top element of stack s
func (s *Stack) Peek() int {
	return s.top.value
}

// RemoveBottom removes bottom element of stack
func (s *Stack) RemoveBottom() int {
	b := s.bottom
	s.bottom = s.bottom.above
	if s.bottom != nil {
		s.bottom.below = nil
	}
	s.size--
	return b.value
}
