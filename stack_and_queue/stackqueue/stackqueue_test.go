package stackqueue

import "testing"

func TestStackQueue(t *testing.T) {
	q := New()

	if s := q.Size(); s != 0 {
		t.Errorf("q.Size() expected 0, got %v", s)
	}
	if v, ok := q.Peek(); ok || v != 0 {
		t.Errorf("q.Peek() expected 0, got %v", v)
	}
	if v, ok := q.Remove(); ok || v != 0 {
		t.Errorf("q.Remove() expected 0, got %v", v)
	}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	if s := q.Size(); s != 3 {
		t.Errorf("q.Size() expected 3, got %v", s)
	}
	if v, ok := q.Peek(); !ok || v != 1 {
		t.Errorf("q.Peek() expected 1, got %v", v)
	}
	if v, ok := q.Remove(); !ok || v != 1 {
		t.Errorf("q.Remove() expected 1, got %v", v)
	}

	if s := q.Size(); s != 2 {
		t.Errorf("q.Size() expected 3, got %v", s)
	}
	if v, ok := q.Peek(); !ok || v != 2 {
		t.Errorf("q.Peek() expected 2, got %v", v)
	}
	if v, ok := q.Remove(); !ok || v != 2 {
		t.Errorf("q.Remove() expected 2, got %v", v)
	}

	if s := q.Size(); s != 1 {
		t.Errorf("q.Size() expected 1, got %v", s)
	}
	if v, ok := q.Peek(); !ok || v != 3 {
		t.Errorf("q.Peek() expected 3, got %v", v)
	}
	if v, ok := q.Remove(); !ok || v != 3 {
		t.Errorf("q.Remove() expected 3, got %v", v)
	}
}
