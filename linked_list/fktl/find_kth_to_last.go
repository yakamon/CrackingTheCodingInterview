package fktl

import (
	"fmt"
)

// PrintKthToLast prints kth to last element of a linked list
func PrintKthToLast(node *Node, k int) int {
	if node == nil {
		return 0
	}
	index := PrintKthToLast(node, k)
	if index == k {
		fmt.Printf("%vth to last node is %v", k, node.Value)
	}
	return index
}

// KthToLast returns node
func KthToLast(n *Node, k int) *Node {
	node1 := n
	node2 := n

	for i := 0; i < k; i++ {
		if node1 == nil {
			return nil
		}
		node1 = node1.next
	}

	for node1 != nil {
		node1 = node1.next
		node2 = node2.next
	}

	return node2
}
