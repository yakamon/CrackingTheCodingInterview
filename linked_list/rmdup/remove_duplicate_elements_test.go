package rmdup

import (
	"container/list"
	"reflect"
	"testing"
)

func makeList(values []int) *list.List {
	l := list.New()
	for _, v := range values {
		l.PushBack(v)
	}
	return l
}

func checkListValues(t *testing.T, l *list.List, values []int) {
	var slice []int
	for e := l.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value.(int))
	}

	// Check values
	if !reflect.DeepEqual(slice, values) {
		t.Errorf("linked list values = %v, expected values = %v", slice, values)
	}
}

func TestRemoveDuplicateElements(t *testing.T) {
	testCase := struct {
		before []int
		after  []int
	}{
		[]int{6, 2, 2, 5, 7, 1, 9, 8, 5, 2},
		[]int{6, 2, 5, 7, 1, 9, 8},
	}

	l := makeList(testCase.before)
	RemoveDuplicateElements(l)
	checkListValues(t, l, testCase.after)
}

func TestRemoveDuplicateElements2(t *testing.T) {
	testCase := struct {
		before []int
		after  []int
	}{
		[]int{6, 2, 2, 5, 7, 1, 9, 8, 5, 2},
		[]int{6, 2, 5, 7, 1, 9, 8},
	}

	l := makeList(testCase.before)
	RemoveDuplicateElements2(l)
	checkListValues(t, l, testCase.after)
}
