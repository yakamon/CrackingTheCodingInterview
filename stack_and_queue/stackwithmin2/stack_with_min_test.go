package stackwithmin2

import "testing"

func TestStack(t *testing.T) {
	s := NewStackWithMin()

	if s.Size() != 0 {
		t.Error("s.Size() should initialized as 0")
	}
	if val, ok := s.Pop(); !ok && val != 0 {
		t.Errorf("Stack should be empty: ok = %v, val = %v", ok, val)
	}
	if val, ok := s.Peek(); !ok && val != 0 {
		t.Errorf("Stack should be empty: ok = %v, val = %v", ok, val)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if val, ok := s.Peek(); ok && val != 3 {
		t.Error("s.Peek() should return 3")
	}
	if size := s.Size(); size != 3 {
		t.Error("s.Size() should return 3")
	}

	if top, ok := s.Pop(); ok && top != 3 {
		t.Errorf("s.Pop() should return 3 here, got %v", top)
	}
	if size := s.Size(); size != 2 {
		t.Errorf("s.Pop() should decremented length 2 here, got %v", size)
	}
	if top, ok := s.Pop(); ok && top != 2 {
		t.Errorf("s.Pop() should return 2 here, got %v", top)
	}
	if size := s.Size(); size != 1 {
		t.Errorf("s.Pop() should decremented length 1 here, got %v", size)
	}
	if top, ok := s.Pop(); ok && top != 1 {
		t.Errorf("s.Pop() should return 1 here, got %v", top)
	}
	if size := s.Size(); size != 0 {
		t.Errorf("s.Pop() should decremented length 0 here, got %v", size)
	}
}

func TestMin(t *testing.T) {
	s := NewStackWithMin()

	s.Push(3)
	if min, ok := s.Min(); ok && min != 3 {
		t.Errorf("want 3, got %v", min)
	}

	s.Push(4)
	if min, ok := s.Min(); ok && min != 3 {
		t.Errorf("want 3, got %v", min)
	}

	s.Push(2)
	if min, ok := s.Min(); ok && min != 2 {
		t.Errorf("want 2, got %v", min)
	}

	s.Pop()
	if min, ok := s.Min(); ok && min != 3 {
		t.Errorf("want 3, got %v", min)
	}

	s.Pop()
	if min, ok := s.Min(); ok && min != 3 {
		t.Errorf("want 3, got %v", min)
	}
}
