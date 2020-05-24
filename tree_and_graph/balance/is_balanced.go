package balance

import "container/list"

// Node is binary tree node.
type Node struct {
	Left  *Node
	Right *Node
}

// NewNode returns a new node of binary tree.
func NewNode() *Node {
	return &Node{nil, nil}
}

// IsBalanced checks if a binary tree is balanced or not.
func IsBalanced(root *Node) bool {
	if root == nil {
		return true
	}

	end := false
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		n := q.Remove(q.Front()).(*Node)
		if end && (n.Left != nil || n.Right != nil) {
			return false
		}
		if n.Left == nil {
			end = true
		} else {
			q.PushBack(n.Left)
		}
		if n.Right == nil {
			end = true
		} else {
			q.PushBack(n.Right)
		}
	}
	return true
}
