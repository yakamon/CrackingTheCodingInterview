package merge

import (
	"crypto/rand"
	"math/big"
	"sort"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	lengths := []int{0, 1, 10, 100, 1000, 10000, 50000, 100000}
	for i, length := range lengths {
		s := makeIntSlice(length)
		Sort(s)
		if !sort.IsSorted(sort.IntSlice(s)) {
			t.Errorf("Case%d: Sort() failed: %v", i, s)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	length := 10000
	for bi := 0; bi < b.N; bi++ {
		s := makeIntSlice(length)
		Sort(s)
	}
}

func makeIntSlice(length int) []int {
	var s []int
	for i := 0; i < length; i++ {
		v, _ := rand.Int(rand.Reader, big.NewInt(1000))
		n, _ := strconv.Atoi(v.String())
		s = append(s, n)
	}
	return s
}
