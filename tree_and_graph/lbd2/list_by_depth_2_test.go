package lbd2

import (
	"container/list"
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
)

func TestCreateLinkedListByIndex(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	l1 := list.New()
	l1.PushBack(New(s[5]))
	l2 := list.New()
	l2.PushBack(New(s[2]))
	l2.PushBack(New(s[8]))
	l3 := list.New()
	l3.PushBack(New(s[1]))
	l3.PushBack(New(s[4]))
	l3.PushBack(New(s[7]))
	l3.PushBack(New(s[9]))
	l4 := list.New()
	l4.PushBack(New(s[0]))
	l4.PushBack(New(s[3]))
	l4.PushBack(New(s[6]))
	testLists := []*list.List{l1, l2, l3, l4}

	tree := makeBST(s)
	lists := MakeLinkedListsForTreeDepth(tree)

	for i := range testLists {
		t.Logf("Depth %v", i)
		l := lists[i]
		tl := testLists[i]
		for l.Len() > 0 && tl.Len() > 0 {
			actual := l.Remove(l.Front())
			expected := tl.Remove(tl.Front())
			if actual.(*Node).Value != expected.(*Node).Value {
				t.Errorf("Want %v, got %v", expected, actual)
			}
		}
		if tl.Len() > 0 {
			t.Error("List lengths is different")
		}
	}
}

func makeBST(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		return New(s[0])
	}

	mid := len(s) / 2
	n := New(s[mid])
	n.Left = makeBST(s[:mid])
	n.Right = makeBST(s[mid+1:])
	return n
}

func makeRandomIntSlice(l int) []int {
	s := make([]int, l, l)
	for i := range s {
		s[i] = generateRandInt()
	}
	return s
}

func generateRandInt() int {
	bi, _ := rand.Int(rand.Reader, big.NewInt(255))
	i, _ := strconv.Atoi(bi.String())
	return i
}
