package search

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
)

const MaxRandomInt = 10000

func TestSearchRotatedArray(t *testing.T) {
	lengths := []int{
		0,
		1,
		5,
		10,
		50,
		100,
		500,
		1000,
		5000,
		10000,
		50000,
		100000,
	}
	for _, length := range lengths {
		slice := createSortAndRotatedIntSlice(length)
		searchIndex, searchValue := createRandomTargetForSearch(slice)
		if ret := SearchRotatedArray(slice, searchValue); ret != searchIndex {
			t.Errorf("SearchRotatedArray() is invalid: target = %v; wanted = %v; got = %v", searchValue, searchIndex, ret)
		}
	}
}

func createSortAndRotatedIntSlice(length int) []int {
	slice := createRandomIntSlice(length)
	heapSort(slice)
	return rotateIntSlice(slice)
}

func createRandomIntSlice(length int) []int {
	slice := make([]int, length, length)
	for i := range slice {
		slice[i] = generateRandomInt(MaxRandomInt)
	}
	return slice
}

func generateRandomInt(max int) int {
	if max <= 1 {
		max = 2
	}
	obj, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	n, _ := strconv.Atoi(obj.String())
	return n
}

func heapSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	buildMaxHeap(slice)
	for last := len(slice) - 1; last >= 0; last-- {
		slice[0], slice[last] = slice[last], slice[0]
		maxHeapify(slice, 0, last-1)
	}
}

func buildMaxHeap(slice []int) {
	last := len(slice) - 1
	for root := (last + 1) / 2; root >= 0; root-- {
		maxHeapify(slice, root, last)
	}
}

func maxHeapify(slice []int, root, last int) {
	v := slice[root]

	i := 2*root + 1
	if i < last && slice[i] < slice[i+1] {
		i++
	}

	for i <= last && v < slice[i] {
		slice[root] = slice[i]
		root = i
		i = 2*root + 1
		if i < last && slice[i] < slice[i+1] {
			i++
		}
	}

	slice[root] = v
}

func rotateIntSlice(slice []int) []int {
	length := len(slice)
	if length <= 1 {
		return slice
	}
	i := generateRandomInt(length-1) + 1
	rotatedSlice := make([]int, length, length)
	copy(rotatedSlice[:length-i], slice[i:])
	copy(rotatedSlice[length-i:], slice[:i])
	return rotatedSlice
}

func createRandomTargetForSearch(slice []int) (int, int) {
	length := len(slice)
	searchIndex := generateRandomInt(2*length - 1)
	if searchIndex >= length {
		return -1, MaxRandomInt + 1
	}
	return searchIndex, slice[searchIndex]
}
