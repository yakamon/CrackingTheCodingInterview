package fib

import "testing"

func provideTestCasesForFib() [][]int {
	return [][]int{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
	}
}

func TestFib1(t *testing.T) {
	t.Log(Fib1(10))
	for _, testCase := range provideTestCasesForFib() {
		n, exp := testCase[0], testCase[1]
		if act := Fib1(n); act != exp {
			t.Errorf("Fib1() failed: want %v, but got %v", exp, act)
		}
	}
}

func TestFib2(t *testing.T) {
	for _, testCase := range provideTestCasesForFib() {
		n, exp := testCase[0], testCase[1]
		if act := Fib2(n); act != exp {
			t.Errorf("Fib1() failed: want %v, but got %v", exp, act)
		}
	}
}

func TestFib3(t *testing.T) {
	for _, testCase := range provideTestCasesForFib() {
		n, exp := testCase[0], testCase[1]
		if act := Fib3(n); act != exp {
			t.Errorf("Fib1() failed: want %v, but got %v", exp, act)
		}
	}
}

func TestFib4(t *testing.T) {
	for _, testCase := range provideTestCasesForFib() {
		n, exp := testCase[0], testCase[1]
		if act := Fib4(n); act != exp {
			t.Errorf("Fib1() failed: want %v, but got %v", exp, act)
		}
	}
}
