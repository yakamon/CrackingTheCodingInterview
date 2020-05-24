package stackqueue

import (
	"../stack"
)

// Queue is queue data structure implemented with two stacks.
type Queue struct {
	new *stack.Stack
	old *stack.Stack
}

// New returns a new Queue instance.
func New() *Queue {
	return &Queue{stack.New(), stack.New()}
}

// Size returns the size of queue
func (q *Queue) Size() int {
	return q.new.Size() + q.old.Size()
}

// Add adds a element to the tail of queue.
func (q *Queue) Add(value int) {
	q.new.Push(value)
}

// Remove removes the first element of queue and returns the value.
func (q *Queue) Remove() (int, bool) {
	q.shiftStacks()
	return q.old.Pop()
}

// Peek returns the first element of queue.
func (q *Queue) Peek() (int, bool) {
	q.shiftStacks()
	return q.old.Peek()
}

func (q *Queue) shiftStacks() {
	if q.old.IsEmpty() {
		for !q.new.IsEmpty() {
			v, _ := q.new.Pop()
			q.old.Push(v)
		}
	}
}
