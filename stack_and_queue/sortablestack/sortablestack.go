package sortstack

import (
	"../stack"
)

// Sort sorts stack
func Sort(s *stack.Stack) {
	reversed := stack.New()
	for v, ok := s.Pop(); ok; v, ok = s.Pop() {
		for peek, ok2 := reversed.Peek(); ok2 && v < peek; peek, ok2 = reversed.Peek() {
			pop, _ := reversed.Pop()
			s.Push(pop)
		}
		reversed.Push(v)
	}

	for v, ok := reversed.Pop(); ok; v, ok = reversed.Pop() {
		s.Push(v)
	}
}
