package merge

import (
	"crypto/rand"
	"math/big"
	"sort"
	"strconv"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	lengths := [][2]int{
		{0, 0},
		{1, 1},
		{1, 2},
		{10, 10},
		{20, 10},
		{100, 200},
		{2000, 1000},
		{10000, 20000},
		{200000, 100000},
	}
	for _, pair := range lengths {
		a, b, al, bl := makeTwoSortedIntSlices(pair[0], pair[1])
		MergeSortedArrays(a, b, al, bl)
		if !sort.IsSorted(sort.IntSlice(a)) {
			t.Error("MergeSortedArrays() faild")
		}
	}
}

func BenchmarkMergeSortedArrays(b *testing.B) {
	lengthPair := [2]int{10000, 20000}
	for bi := 0; bi < b.N; bi++ {
		a, b, al, bl := makeTwoSortedIntSlices(lengthPair[0], lengthPair[1])
		MergeSortedArrays(a, b, al, bl)
	}
}

func makeTwoSortedIntSlices(l1, l2 int) ([]int, []int, int, int) {
	lenSum := l1 + l2
	a, b := make([]int, lenSum, lenSum), make([]int, l2, l2)
	for i := 0; i < l1; i++ {
		a[i] = randInt()
	}
	for i := 0; i < l2; i++ {
		b[i] = randInt()
	}
	quickSort(a)
	quickSort(b)
	return a, b, l1, l2
}

func randInt() int {
	obj, _ := rand.Int(rand.Reader, big.NewInt(100000))
	n, _ := strconv.Atoi(obj.String())
	return n
}

func makeIntSlice(length int) []int {
	s := make([]int, length, length)
	for i := 0; i < length; i++ {
		s[i] = randInt()
	}
	return s
}

func quickSort(s []int) {
	quickSortInner(s, 0, len(s)-1)
}

func quickSortInner(s []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(s, l, r)
	if l < i-1 {
		quickSortInner(s, l, i-1)
	}
	if i < r {
		quickSortInner(s, i, r)
	}
}

func partition(s []int, l, r int) int {
	pivot := s[(l+r)/2]
	for l <= r {
		for s[l] < pivot {
			l++
		}
		for pivot < s[r] {
			r--
		}
		if l <= r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}
	return l
}

func TestSort(t *testing.T) {
	lengths := []int{0, 1, 10, 100, 1000}
	for _, length := range lengths {
		s := makeIntSlice(length)
		quickSort(s)
		if !sort.IsSorted(sort.IntSlice(s)) {
			t.Errorf("quickSort() faild: %v", s)
		}
	}
}
