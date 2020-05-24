package magicindex

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
)

func BenchmarkFindMagicIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(FindMagicIndex(makeRandomIntSlice(100)))
	}
}

func BenchmarkFindMagicIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(FindMagicIndex2(makeRandomIntSlice(100)))
	}
}

func makeRandomIntSlice(size int) []int {
	s := make([]int, size, size)
	for i := range s {
		v, _ := rand.Int(rand.Reader, big.NewInt(1000))
		n, _ := strconv.Atoi(v.String())
		s[i] = n
	}
	return s
}
