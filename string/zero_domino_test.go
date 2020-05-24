package chapter01

import (
	"reflect"
	"testing"
)

func TestZeroDomino(t *testing.T) {
	testCaseZeroDomino := []struct {
		In  [][]int
		Out [][]int
	}{
		{
			[][]int{
				{1, 2, 0},
				{0, 4, 5},
				{10, 2, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 2, 0},
			}},
		{
			[][]int{
				{1, 4, 5, 6, 9, 10, 45, 0},
				{52, 2, 6, 2, 6, 2, 6, 60},
			},
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{52, 2, 6, 2, 6, 2, 6, 0},
			},
		},
	}

	for i, test := range testCaseZeroDomino {
		t.Logf("TestCase %d", i)
		if actual := ZeroDomino(test.In); !reflect.DeepEqual(actual, test.Out) {
			t.Errorf("  Expected: %v\n", test.Out)
			t.Errorf("  Got: %v\n", actual)
		}
	}
}

func BenchmarkZeroDomino(b *testing.B) {
	testCaseZeroDomino := []struct {
		In  [][]int
		Out [][]int
	}{
		{
			[][]int{
				{1, 2, 0},
				{0, 4, 5},
				{10, 2, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 2, 0},
			}},
		{
			[][]int{
				{1, 4, 5, 6, 9, 10, 45, 0},
				{52, 2, 6, 2, 6, 2, 6, 60},
			},
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{52, 2, 6, 2, 6, 2, 6, 0},
			},
		},
	}
	for iter := 0; iter < b.N; iter++ {
		for _, test := range testCaseZeroDomino {
			b.Log(ZeroDomino(test.In))
		}
	}
}

func TestZeroDomino2(t *testing.T) {
	testCaseZeroDomino := []struct {
		In  [][]int
		Out [][]int
	}{
		{
			[][]int{
				{1, 2, 0},
				{0, 4, 5},
				{10, 2, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 2, 0},
			}},
		{
			[][]int{
				{1, 4, 5, 6, 9, 10, 45, 0},
				{52, 2, 6, 2, 6, 2, 6, 60},
			},
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{52, 2, 6, 2, 6, 2, 6, 0},
			},
		},
	}
	for i, test := range testCaseZeroDomino {
		t.Logf("TestCase %d", i)
		if actual := ZeroDomino2(test.In); !reflect.DeepEqual(actual, test.Out) {
			t.Errorf("  Expected: %v\n", test.Out)
			t.Errorf("  Got: %v\n", actual)
		}
	}
}

func BenchmarkZeroDomino2(b *testing.B) {
	testCaseZeroDomino := []struct {
		In  [][]int
		Out [][]int
	}{
		{
			[][]int{
				{1, 2, 0},
				{0, 4, 5},
				{10, 2, 9},
			},
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 2, 0},
			}},
		{
			[][]int{
				{1, 4, 5, 6, 9, 10, 45, 0},
				{52, 2, 6, 2, 6, 2, 6, 60},
			},
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{52, 2, 6, 2, 6, 2, 6, 0},
			},
		},
	}
	for iter := 0; iter < b.N; iter++ {
		for _, test := range testCaseZeroDomino {
			b.Log(ZeroDomino2(test.In))
		}
	}
}
