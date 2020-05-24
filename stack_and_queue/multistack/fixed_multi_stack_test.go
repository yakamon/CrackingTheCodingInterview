package multistack

import "testing"

func TestFixedMultiStack(t *testing.T) {
	s := NewFixedMultiStack(2)

	if size := s.Size(0); size != 0 {
		t.Errorf("want 0, got %v", size)
	}
	if peek := s.Peek(0); peek != nil {
		t.Errorf("want 0, got %v", peek)
	}
	if topIndex := s.IndexOfTop(0); topIndex != -1 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if err := s.Push(0, 1); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(0); topIndex != 0 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if s.IsFull(0) {
		t.Error("Stack should not be full yet")
	}
	if err := s.Push(0, 4); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(0); topIndex != 1 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if !s.IsFull(0) {
		t.Error("Stack should not be full yet")
	}
	if pop := s.Pop(0); pop != 4 {
		t.Errorf("want 4, got %v", pop)
	}

	if size := s.Size(1); size != 0 {
		t.Errorf("want 0, got %v", size)
	}
	if peek := s.Peek(1); peek != nil {
		t.Errorf("want 0, got %v", peek)
	}
	if topIndex := s.IndexOfTop(1); topIndex != -1 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if err := s.Push(1, 2); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(1); topIndex != 2 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if s.IsFull(1) {
		t.Error("Stack should not be full yet")
	}
	if err := s.Push(1, 5); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(1); topIndex != 3 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if !s.IsFull(1) {
		t.Error("Stack should not be full yet")
	}
	if pop := s.Pop(1); pop != 5 {
		t.Errorf("want 5, got %v", pop)
	}

	if size := s.Size(2); size != 0 {
		t.Errorf("want 0, got %v", size)
	}
	if peek := s.Peek(2); peek != nil {
		t.Errorf("want 0, got %v", peek)
	}
	if topIndex := s.IndexOfTop(2); topIndex != -1 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if err := s.Push(2, 3); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(2); topIndex != 4 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if s.IsFull(2) {
		t.Error("Stack should not be full yet")
	}
	if err := s.Push(2, 6); err != nil {
		t.Error("No error should occur here")
	}
	if topIndex := s.IndexOfTop(2); topIndex != 5 {
		t.Errorf("want -1, got %v", topIndex)
	}
	if !s.IsFull(2) {
		t.Error("Stack should not be full yet")
	}
	if pop := s.Pop(2); pop != 6 {
		t.Errorf("want 6, got %v", pop)
	}
}
