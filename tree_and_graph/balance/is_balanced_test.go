package balance

import (
	"testing"
)

func makeBalancedBinaryTree(depth int) *Node {
	if depth == 0 {
		return nil
	}
	root := NewNode()
	root.Left = makeBalancedBinaryTree(depth - 1)
	root.Right = makeBalancedBinaryTree(depth - 1)
	return root
}

func TestIsBalanced(t *testing.T) {
	balanced := makeBalancedBinaryTree(5)
	if isBalanced := IsBalanced(balanced); !isBalanced {
		t.Errorf("Expected: %v, got: %v", true, isBalanced)
	}

	unbalanced := makeBalancedBinaryTree(5)
	n := NewNode()
	n.Left = NewNode()
	unbalanced.Left.Left.Left.Left.Left = n
	if isBalanced := IsBalanced(unbalanced); isBalanced {
		t.Errorf("Expected: %v, got: %v", false, isBalanced)
	}
}
