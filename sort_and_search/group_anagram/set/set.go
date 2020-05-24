package set

import (
	"container/list"
)

// Set is set data structure.
type Set struct {
	d *list.List
	m map[interface{}]*list.Element
}

// New initializes new set.
func New(slice []interface{}) *Set {
	set := &Set{list.New(), map[interface{}]*list.Element{}}
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

// Len returns length of set.
func (s *Set) Len() int {
	return s.d.Len()
}

// Add adds element to set.
func (s *Set) Add(x interface{}) {
	if _, ok := s.m[x]; ok {
		return
	}
	s.m[x] = s.d.PushBack(x)
}

// Remove removes element from set.
func (s *Set) Remove(x interface{}) {
	if _, ok := s.m[x]; !ok {
		return
	}
	s.d.Remove(s.m[x])
	delete(s.m, x)
}

// Union returns sum set of s and a.
func (s *Set) Union(a *Set) *Set {
	union := New(nil)
	for e := s.d.Front(); e != nil; e = e.Next() {
		union.Add(e.Value)
	}
	for e := a.d.Front(); e != nil; e = e.Next() {
		union.Add(e.Value)
	}
	return union
}

// Intersection returns intersection set of s and a.
func (s *Set) Intersection(a *Set) *Set {
	union := s.Union(a)
	diff1 := s.Diff(a)
	diff2 := a.Diff(s)
	return union.Diff(diff1).Diff(diff2)
}

// Diff returns difference set of s and a.
func (s *Set) Diff(a *Set) *Set {
	diff := New(nil)
	for e := s.d.Front(); e != nil; e = e.Next() {
		diff.Add(e.Value)
	}
	for e := a.d.Front(); e != nil; e = e.Next() {
		diff.Remove(e.Value)
	}
	return diff
}

// SymDiff returns symmetric difference set of s and a.
func (s *Set) SymDiff(a *Set) *Set {
	diff1 := s.Diff(a)
	diff2 := a.Diff(s)
	return diff1.Union(diff2)
}

// Equal checks if s and a have same elements.
func (s *Set) Equal(a *Set) bool {
	return s.IsSubsetOf(a) && s.IsSupersetOf(a)
}

// IsSubsetOf checks if s is subset of a.
func (s *Set) IsSubsetOf(a *Set) bool {
	return s.Diff(a).Len() == 0
}

// IsSupersetOf checks if s is superset of a.
func (s *Set) IsSupersetOf(a *Set) bool {
	return a.Diff(s).Len() == 0
}

// Slice returns elements of set as slice.
func (s *Set) Slice() []interface{} {
	slice := make([]interface{}, 0, s.Len())
	for e := s.d.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value)
	}
	return slice
}
