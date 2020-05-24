package powerset

import (
	"testing"
)

func BenchmarkPowerSet(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for i := 0; i < 20; i++ {
			set := makeSet(i)
			powerSet := PowerSet(set)
			b.Log(len(powerSet))
		}
	}
}

func BenchmarkPowerSet2(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for i := 0; i < 20; i++ {
			set := makeSet(i)
			powerSet := PowerSet2(set)
			b.Log(len(powerSet))
		}
	}
}

func BenchmarkPowerSet3(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for i := 0; i < 20; i++ {
			set := makeSet(i)
			powerSet := PowerSet3(set)
			b.Log(len(powerSet))
		}
	}
}

func BenchmarkPowerSet4(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for i := 0; i < 20; i++ {
			set := makeSet(i)
			powerSet := PowerSet4(set)
			b.Log(len(powerSet))
		}
	}
}

func BenchmarkPowerSet5(b *testing.B) {
	for x := 0; x < b.N; x++ {
		for i := 0; i < 20; i++ {
			set := makeSet(i)
			powerSet := PowerSet5(set)
			b.Log(len(powerSet))
		}
	}
}

func makeSet(size int) []int {
	set := make([]int, size, size)
	for i := 0; i < size; i++ {
		set[i] = i + 1
	}
	return set
}
