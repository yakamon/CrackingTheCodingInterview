package lbd2

import (
	"container/list"
)

// Node is tree node.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// New returns a new node of tree.
func New(v int) *Node {
	return &Node{v, nil, nil}
}

// MakeLinkedListsForTreeDepth make list of linked lists that have elements in each tree depth.
func MakeLinkedListsForTreeDepth(root *Node) []*list.List {
	if root == nil {
		return nil
	}

	var lists []*list.List
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		lists = append(lists, l)
		parents := copyList(l)
		l = list.New()
		for parents.Len() > 0 {
			n := parents.Remove(parents.Front())
			if left := n.(*Node).Left; left != nil {
				l.PushBack(left)
			}
			if right := n.(*Node).Right; right != nil {
				l.PushBack(right)
			}
		}
	}

	return lists
}

func copyList(l *list.List) *list.List {
	cp := list.New()
	for e := l.Front(); e != nil; e = e.Next() {
		cp.PushBack(e.Value)
	}
	return cp
}
