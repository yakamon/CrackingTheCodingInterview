package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()

	if s := q.Size(); s != 0 {
		t.Errorf("Queue.Size() expected 0, got %v", s)
	}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	if s := q.Size(); s != 3 {
		t.Errorf("Queue.Size() expected 3, got %v", s)
	}
	if v, ok := q.Peek(); !ok || v != 1 {
		t.Errorf("Queue.Peek() expected 1, got %v", v)
	}

	if v, ok := q.Remove(); !ok || v != 1 {
		t.Errorf("Queue.Remove() expected 1, got %v", v)
	}
	if s := q.Size(); s != 2 {
		t.Errorf("Queue.Size() expected 2, got %v", s)
	}

	if v, ok := q.Remove(); !ok || v != 2 {
		t.Errorf("Queue.Remove() expected 2, got %v", v)
	}
	if s := q.Size(); s != 1 {
		t.Errorf("Queue.Size() expected 1, got %v", s)
	}

	if v, ok := q.Remove(); !ok || v != 3 {
		t.Errorf("Queue.Remove() expected 3, got %v", v)
	}
	if s := q.Size(); s != 0 {
		t.Errorf("Queue.Size() expected 0, got %v", s)
	}
}
