package stackset

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(4)

	if s.Size() != 0 {
		t.Error("s.Size should initialized as 0")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if top := s.Peek(); top != 3 {
		t.Error("s.Peek() should return 3")
	}
	if s.Size() != 3 {
		t.Error("s.Size() should return 3")
	}

	if top := s.Pop(); top != 3 {
		t.Errorf("s.Pop() should return 3 here, got %v", top)
	}
	if size := s.Size(); size != 2 {
		t.Errorf("s.Pop() should decremented size 2 here, got %v", size)
	}
	if top := s.Pop(); top != 2 {
		t.Errorf("s.Pop() should return 2 here, got %v", top)
	}
	if size := s.Size(); size != 1 {
		t.Errorf("s.Pop() should decremented size 1 here, got %v", size)
	}
	if top := s.Pop(); top != 1 {
		t.Errorf("s.Pop() should return 1 here, got %v", top)
	}
	if size := s.Size(); size != 0 {
		t.Errorf("s.Pop() should decremented size 0 here, got %v", size)
	}
}

func TestStackSet(t *testing.T) {
	percap := 2
	s := NewStackSet(percap)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)

	if n := s.NumOfStacks(); n != 3 {
		t.Errorf("s.NumOfStacks() expected 3, got %v", n)
	}
	if v, ok := s.Pop(); !ok || v != 6 {
		t.Errorf("s.Pop() expected 6, got %v", v)
	}

	if n := s.NumOfStacks(); n != 3 {
		t.Errorf("s.NumOfStacks() expected 3, got %v", n)
	}
	if v, ok := s.Pop(); !ok || v != 5 {
		t.Errorf("s.Pop() expected 5, got %v", v)
	}

	if v, ok := s.Pop(); !ok || v != 4 {
		t.Errorf("s.Pop() expected 4, got %v", v)
	}

	if v, ok := s.Pop(); !ok || v != 3 {
		t.Errorf("s.Pop() expected 3, got %v", v)
	}

	if v, ok := s.Pop(); !ok || v != 2 {
		t.Errorf("s.Pop() expected 2, got %v", v)
	}

	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("s.Pop() expected 1, got %v", v)
	}
}

func TestStackSetPopAt(t *testing.T) {
	percap := 2
	s := NewStackSet(percap)

	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Push(13)
	s.Push(14)
	s.Push(15)
	s.Push(16)

	if v, ok := s.PopAt(0); !ok || v != 11 {
		t.Errorf("s.Pop() expected 11, got %v", v)
	}
	if v, ok := s.PopAt(1); !ok || v != 14 {
		t.Errorf("s.Pop() expected 14, got %v", v)
	}
	if v, ok := s.PopAt(2); !ok || v != 16 {
		t.Errorf("s.Pop() expected 11, got %v", v)
	}
	if n := s.NumOfStacks(); n != 2 {
		t.Errorf("s.NumOfStacks() expected 2, got %v", n)
	}
}
