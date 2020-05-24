package intmul

import "testing"

func BenchmarkMul(b *testing.B) {
	for k := 0; k < b.N; k++ {
		for i := 100; i < 200000; i++ {
			b.Log(Mul(1000, i+1))
		}
	}
}

func BenchmarkMul2(b *testing.B) {
	for k := 0; k < b.N; k++ {
		for i := 100; i < 200000; i++ {
			b.Log(Mul2(1000, i+1))
		}
	}
}

func BenchmarkMul3(b *testing.B) {
	for k := 0; k < b.N; k++ {
		for i := 100; i < 200000; i++ {
			b.Log(Mul3(1000, i+1))
		}
	}
}
