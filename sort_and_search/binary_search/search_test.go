package search

import (
	"math/rand"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	lengths := []int{
		0,
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
	}
	for _, length := range lengths {
		s := makeIntSlice(length)
		mergeSort(s)
		if !sort.IsSorted(sort.IntSlice(s)) {
			t.Error("mergeSort() result is invalid.")
		}
		i, v := getRandomIndexAndValue(s)
		if result := BinarySearch(s, v); result != i {
			t.Errorf("BinarySearch() failed: got %v, expected %v", result, i)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	length := 10000
	for bi := 0; bi < b.N; bi++ {
		s := makeIntSlice(length)
		mergeSort(s)
		if !sort.IsSorted(sort.IntSlice(s)) {
			b.Error("mergeSort() result is invalid.")
		}
		_, v := getRandomIndexAndValue(s)
		BinarySearch(s, v)
	}
}

func makeIntSlice(length int) []int {
	s := make([]int, length, length)
	m := map[int]bool{}
	i := 0
	for i < len(s) {
		n := rand.Intn(10000000)
		if exists, _ := m[n]; !exists {
			m[n] = true
			s[i] = n
			i++
		}
	}
	return s
}

func mergeSort(s []int) {
	h := make([]int, len(s), len(s))
	mergeSortInner(s, h, 0, len(s)-1)
}

func mergeSortInner(s, h []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	mergeSortInner(s, h, l, m)
	mergeSortInner(s, h, m+1, r)
	merge(s, h, l, m, r)
}

func merge(s, h []int, l, m, r int) {
	copy(h[l:r+1], s[l:r+1])

	li, ri := l, m+1
	cur := l
	for li <= m && ri <= r {
		if h[li] < h[ri] {
			s[cur] = h[li]
			li++
		} else {
			s[cur] = h[ri]
			ri++
		}
		cur++
	}

	if li <= m {
		copy(s[cur:], h[li:m+1])
	} else if ri <= r {
		copy(s[cur:], h[ri:r+1])
	}
}

func getRandomIndexAndValue(s []int) (int, int) {
	if len(s) == 0 {
		return -1, 0
	}
	if len(s) == 1 {
		return 0, s[0]
	}
	i := rand.Intn(len(s) - 1)
	return i, s[i]
}
