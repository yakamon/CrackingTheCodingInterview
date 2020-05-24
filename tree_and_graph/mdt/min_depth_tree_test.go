package mt

import (
	"crypto/rand"
	"math/big"
	"sort"
	"strconv"
	"testing"
)

func TestCreateBSTFromSortedSlice(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := createSortedRandomIntSlice(1000)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		tree := CreateBSTFromSortedSlice(s)
		if !isValidBST(tree) {
			t.Errorf("CreateBSTFromSortedSlice is not a valid BST.")
		}
	}
}

func createSortedRandomIntSlice(l int) []int {
	s := make([]int, l, l)
	for i := range s {
		v, _ := rand.Int(rand.Reader, big.NewInt(1000))
		n, _ := strconv.Atoi(v.String())
		s[i] = n
	}
	return s
}

func isValidBST(root *Node) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Value < root.Left.Value) || (root.Right != nil && root.Value > root.Right.Value) {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
