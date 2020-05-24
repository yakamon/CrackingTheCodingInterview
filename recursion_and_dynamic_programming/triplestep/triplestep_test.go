package triplestep

import "testing"

func provideTestCasesForTripleStep() [][]int {
	return [][]int{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
	}
}

func TestTripleStep(t *testing.T) {
	for _, c := range provideTestCasesForTripleStep() {
		n, e := c[0], c[1]
		if a := TripleStep(n); a != e {
			t.Errorf("TripleStep(%v) failed: want %v, got %v", n, e, a)
		}
	}
}

func TestTripleStep2(t *testing.T) {
	for _, c := range provideTestCasesForTripleStep() {
		n, e := c[0], c[1]
		if a := TripleStep2(n); a != e {
			t.Errorf("TripleStep(%v) failed: want %v, got %v", n, e, a)
		}
	}
}

func BenchmarkTripleStep2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < 1000; n++ {
			b.Logf("TripleStep2(%v) = %v", n, TripleStep2(n))
		}
	}
}
