package multistack

import (
	"fmt"
)

var numberOfStacks = 3

// FixedMultiStack is a data structure consisting of three stacks implemented in one array
type FixedMultiStack struct {
	stackCap int
	values   []interface{}
	sizes    []int
}

// NewFixedMultiStack creates new instance of FixedMultiStack
func NewFixedMultiStack(stackSize int) *FixedMultiStack {
	values := make([]interface{}, stackSize*numberOfStacks)
	sizes := make([]int, numberOfStacks)

	return &FixedMultiStack{stackSize, values, sizes}
}

// Push adds a new element to nth stack
func (s *FixedMultiStack) Push(index int, value interface{}) error {
	if s.IsFull(index) {
		return fmt.Errorf("%vth stack is full", index)
	}

	s.sizes[index]++
	s.values[s.IndexOfTop(index)] = value
	return nil
}

// Pop pops an element off the stack
func (s *FixedMultiStack) Pop(index int) interface{} {
	if s.IsEmpty(index) {
		return nil
	}

	topIndex := s.IndexOfTop(index)
	topValue := s.values[topIndex]
	s.values[topIndex] = 0
	s.sizes[index]--
	return topValue
}

// Peek returns the top element of nth stack
func (s *FixedMultiStack) Peek(index int) interface{} {
	if s.IsEmpty(index) {
		return nil
	}

	return s.values[s.IndexOfTop(index)]
}

// Size returns the size of nth stack
func (s *FixedMultiStack) Size(index int) int {
	return s.sizes[index]
}

// IsEmpty checks if nth stack is empty or not
func (s *FixedMultiStack) IsEmpty(index int) bool {
	return s.sizes[index] == 0
}

// IsFull checks if nth stack is full or not
func (s *FixedMultiStack) IsFull(index int) bool {
	return s.sizes[index] == s.stackCap
}

// IndexOfTop returns index of top element of nth stack
func (s *FixedMultiStack) IndexOfTop(index int) int {
	offset := index * s.stackCap
	size := s.sizes[index]
	if size == 0 {
		return -1
	}
	return offset + size - 1
}
