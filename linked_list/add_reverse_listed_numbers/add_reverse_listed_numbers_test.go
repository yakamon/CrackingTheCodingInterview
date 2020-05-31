package numadd

import (
	"reflect"
	"testing"
)

var testData = [][3][]interface{}{
	{
		{1, 2, 3}, // linked list A
		{4, 5, 6}, // linked list B
		{5, 7, 9}, // result
	},
	{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{6, 8, 0, 3, 1},
	},
}

func TestAddReverseListedNumbers(t *testing.T) {
	for _, td := range testData {
		a := NewListFromSlice(td[0])
		b := NewListFromSlice(td[1])
		sum := AddReverseListedNumbers(a, b)
		actual := sum.Values()
		expect := td[2]
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("The result of the addition is incorrect.\nExpected: %v\nActual: %v\n", expect, actual)
		}
	}
}

func BenchmarkAddReverseListedNumbers(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, td := range testData {
			l1 := NewListFromSlice(td[0])
			l2 := NewListFromSlice(td[1])
			b.Log(AddReverseListedNumbers(l1, l2))
		}
	}
}

func TestAddReverseListedNumber2(t *testing.T) {
	for _, td := range testData {
		a := NewListFromSlice(td[0])
		b := NewListFromSlice(td[1])
		sum := AddReverseListedNumbers2(a, b)
		actual := sum.Values()
		expect := td[2]
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("The result of the addition is incorrect.\nExpected: %v\nActual: %v\n", expect, actual)
		}
	}
}

func BenchmarkAddReverseListedNumbers2(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, td := range testData {
			l1 := NewListFromSlice(td[0])
			l2 := NewListFromSlice(td[1])
			b.Log(AddReverseListedNumbers2(l1, l2))
		}
	}
}
