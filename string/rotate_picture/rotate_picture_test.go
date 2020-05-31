package chapter01

import (
	"reflect"
	"testing"
)

var testCaseRotatePicture = []struct {
	Case     struct{ Picture [][]int }
	Expected [][]int
}{
	{
		struct{ Picture [][]int }{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}},
		[][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		},
	},
	{
		struct{ Picture [][]int }{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}},
		[][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		},
	},
}

func TestRotatePicture(t *testing.T) {
	for i, test := range testCaseRotatePicture {
		t.Logf("TestCase %d", i)
		if actual := RotatePicture(test.Case.Picture); !reflect.DeepEqual(actual, test.Expected) {
			t.Errorf("  Expected: %v\n", test.Expected)
			t.Errorf("  Got: %v\n", actual)
		}
	}
}

func BenchmarkRotatePicture(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range testCaseRotatePicture {
			b.Log(RotatePicture(test.Case.Picture))
		}
	}
}

func TestRotatePicture2(t *testing.T) {
	for i, test := range testCaseRotatePicture {
		t.Logf("TestCase %d", i)
		if actual := RotatePicture2(test.Case.Picture); !reflect.DeepEqual(actual, test.Expected) {
			t.Errorf("  Expected: %v\n", test.Expected)
			t.Errorf("  Got: %v\n", actual)
		}
	}
}

func BenchmarkRotatePicture2(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range testCaseRotatePicture {
			b.Log(RotatePicture2(test.Case.Picture))
		}
	}
}
