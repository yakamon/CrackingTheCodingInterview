package fktl

import (
	"fmt"
	"testing"
)

func printAllListValues(n *Node) {
	i := 0
	node := n
	for node != nil {
		fmt.Printf("%vth value is: %v\n", i, node.Value)
		node = node.next
		i++
	}
}

func initListForTest() *Node {
	td := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	node := NewNode(td[0])
	for _, v := range td[1:] {
		node.AppendToTail(v)
	}

	return node
}

func TestKthToLast(t *testing.T) {
	for k, value := range map[int]int{1: 10, 2: 9, 4: 7} {
		node := initListForTest()
		n := KthToLast(node, k)
		if n.Value != value {
			t.Errorf("Actual: %v, expected: %v", n.Value, value)
		}
	}
}
