package lbd

import "container/list"

// Node is BST node.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// New returns a new BST node.
func New(v int) *Node {
	return &Node{v, nil, nil}
}

// MakeLinkedListsForTreeDepth makes a slice of linked list of nodes for each tree depth.
func MakeLinkedListsForTreeDepth(root *Node) []*list.List {
	var lists []*list.List
	lists = makeLinkedListsForTreeDepth(root, lists, 0)
	return lists
}

func makeLinkedListsForTreeDepth(root *Node, lists []*list.List, level int) []*list.List {
	if root == nil {
		return lists
	}

	if len(lists) == level {
		lists = append(lists, list.New())
	}
	l := lists[level]
	l.PushBack(root)

	lists = makeLinkedListsForTreeDepth(root.Left, lists, level+1)
	lists = makeLinkedListsForTreeDepth(root.Right, lists, level+1)

	return lists
}
