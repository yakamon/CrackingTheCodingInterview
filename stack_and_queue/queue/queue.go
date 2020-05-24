package queue

// Queue is a queue data structure.
type Queue struct {
	first *node
	last  *node
	size  int
}

type node struct {
	value int
	next  *node
}

// New returns a new queue instance.
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Size returns the size of queue.
func (q *Queue) Size() int {
	return q.size
}

// IsEmpty checks if queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Add adds a new element to the last of queue.
func (q *Queue) Add(value int) {
	n := &node{value, nil}
	if q.last != nil {
		q.last.next = n
	}
	q.last = n
	if q.first == nil {
		q.first = q.last
	}
	q.size++
}

// Remove removes the first element of queue q and returns the value of it.
func (q *Queue) Remove() (int, bool) {
	if q.size == 0 {
		return 0, false
	}

	n := q.first
	q.first = n.next
	if q.first == nil {
		q.last = nil
	}
	q.size--
	return n.value, true
}

// Peek returns the value of first element of queue.
func (q *Queue) Peek() (int, bool) {
	if q.size == 0 {
		return 0, false
	}

	return q.first.value, true
}
