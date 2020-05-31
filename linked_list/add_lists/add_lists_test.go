package numadd2

import (
	"reflect"
	"testing"
)

var testTable = [][3][]int{
	{
		{1, 2, 3},
		{4, 5, 6},
		{5, 7, 9},
	},
	{
		{1, 2, 3, 4},
		{5, 6, 7},
		{1, 8, 0, 1},
	},
}

func TestAddLists(t *testing.T) {
	for _, td := range testTable {
		input1, input2, expected := td[0], td[1], td[2]
		l1, l2 := NewListFromSlice(input1), NewListFromSlice(input2)
		if actual := AddLists(l1, l2).Values(); !reflect.DeepEqual(actual, expected) {
			t.Errorf("The result of addition is incorrect.\nExpected: %v\nActual:   %v", expected, actual)
		}
	}
}

func BenchmarkAddLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testTable {
			input1, input2 := td[0], td[1]
			l1, l2 := NewListFromSlice(input1), NewListFromSlice(input2)
			b.Log(AddLists(l1, l2))
		}
	}
}
