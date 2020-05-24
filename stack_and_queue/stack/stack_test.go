package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()

	if s := s.Size(); s != 0 {
		t.Errorf("s.Size() expected 0, got %v", s)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if v, ok := s.Peek(); !ok || v != 3 {
		t.Errorf("s.Peek() expected 3, got %v", v)
	}
	if s := s.Size(); s != 3 {
		t.Errorf("s.Size() expected 0, got %v", s)
	}

	if v, ok := s.Pop(); !ok || v != 3 {
		t.Errorf("s.Peek() expected 3, got %v", v)
	}
	if s := s.Size(); s != 2 {
		t.Errorf("s.Size() expected 0, got %v", s)
	}
	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("s.Peek() expected 3, got %v", v)
	}
	if s := s.Size(); s != 1 {
		t.Errorf("s.Size() expected 0, got %v", s)
	}
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("s.Peek() expected 3, got %v", v)
	}
	if s := s.Size(); s != 0 {
		t.Errorf("s.Size() expected 0, got %v", s)
	}
}
