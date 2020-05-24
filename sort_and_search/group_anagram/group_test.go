package anagram

import (
	"crypto/rand"
	"math/big"
	"sort"
	"strconv"
	"testing"
)

func TestHeapSort(t *testing.T) {
	lengths := []int{
		0,
		1,
		10,
		100,
		1000,
		10000,
		100000,
	}
	for _, length := range lengths {
		s := makeRuneSlice(length)
		heapSort(s)
		if !sort.SliceIsSorted(s, func(i, j int) bool { return s[i] < s[j] }) {
			t.Error("heapSort() is invalid")
		}
	}
}

func BenchmarkHeapSort(b *testing.B) {
	length := 10000
	for bi := 0; bi < b.N; bi++ {
		s := makeRuneSlice(length)
		heapSort(s)
	}
}

func makeRuneSlice(length int) []rune {
	s := make([]rune, length, length)
	for i := range s {
		s[i] = getRandomAlphabet()
	}
	return s
}

func getRandomAlphabet() rune {
	v, _ := rand.Int(rand.Reader, big.NewInt(int64('z')-int64('a')+1))
	n, _ := strconv.Atoi(v.String())
	return rune(n + int('a'))
}
