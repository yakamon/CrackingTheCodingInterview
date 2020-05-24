package sortstack

import (
	"testing"

	"../stack"
)

func TestSortStack(t *testing.T) {
	s := stack.New()
	for _, v := range []int{2, 3, 5, 1, 9, 0, 4, 6, 8, 7} {
		s.Push(v)
	}

	Sort(s)

	for i, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if pop, ok := s.Pop(); !ok || pop != v {
			t.Errorf("s.Sort() failed: %vth element: expected %v, got %v", i, v, pop)
		}
	}
}
