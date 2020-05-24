package set

import (
	"testing"
)

func TestSetAddRemove(t *testing.T) {
	a := New(nil)
	if ret := a.Len(); ret != 0 {
		t.Errorf("Len() is invalid: %v", ret)
	}

	a.Add(3)
	if ret := a.Len(); ret != 1 {
		t.Errorf("Len() is invalid: %v", ret)
	}
	if !a.Equal(New([]interface{}{3})) {
		t.Errorf("Remove() is invalid: %v", a.Slice())
	}

	a.Add(1)
	if ret := a.Len(); ret != 2 {
		t.Errorf("Len() is invalid: %v", ret)
	}
	if !a.Equal(New([]interface{}{3, 1})) {
		t.Errorf("Remove() is invalid: %v", a.Slice())
	}

	a.Add(2)
	if ret := a.Len(); ret != 3 {
		t.Errorf("Len() is invalid: %v", ret)
	}
	if !a.Equal(New([]interface{}{3, 1, 2})) {
		t.Errorf("Remove() is invalid: %v", a.Slice())
	}

	a.Remove(1)
	if ret := a.Len(); ret != 2 {
		t.Errorf("Len() is invalid: %v", ret)
	}
	if !a.Equal(New([]interface{}{3, 2})) {
		t.Errorf("Remove() is invalid: %v", a.Slice())
	}
}

func TestSetOperations(t *testing.T) {
	sa := []interface{}{1, 2, 3, 4}          // base set
	sb := []interface{}{1, 2, 3, 4}          // equal to a
	sc := []interface{}{2, 3, 4, 5}          // for equal, union, intersection, diff, symdiff
	sd := []interface{}{2, 3}                // subset of a
	se := []interface{}{0, 1, 2, 3, 4, 5, 6} // superset of a

	a := New(sa)
	b := New(sb)
	c := New(sc)
	d := New(sd)
	e := New(se)

	if ret := a.Equal(b); !ret {
		t.Errorf("Equal() is invalid: %v", ret)
	}
	if ret := a.Equal(c); ret {
		t.Errorf("Equal() is invalid: %v", ret)
	}

	acUnion := New([]interface{}{1, 2, 3, 4, 5})
	if union := a.Union(c); !union.Equal(acUnion) {
		t.Errorf("Union() is invalid: %v", union.m)
	}
	acIntersection := New([]interface{}{2, 3, 4})
	if intersection := a.Intersection(c); !intersection.Equal(acIntersection) {
		t.Errorf("Intersection() is invalid: %v", intersection.m)
	}
	acDiff := New([]interface{}{1})
	if diff := a.Diff(c); !diff.Equal(acDiff) {
		t.Errorf("Union() is invalid: %v", diff.m)
	}
	acSymDiff := New([]interface{}{1, 5})
	if symDiff := a.SymDiff(c); !symDiff.Equal(acSymDiff) {
		t.Errorf("Union() is invalid: %v", symDiff.m)
	}

	if ret := d.IsSubsetOf(a); !ret {
		t.Errorf("IsSubsetOf() is invalid: %v", ret)
	}
	if ret := e.IsSupersetOf(a); !ret {
		t.Errorf("IsSupersetOf() is invalid: %v", ret)
	}
}
