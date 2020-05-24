package fktl

// Node is node of simple linked list
type Node struct {
	next  *Node
	Value interface{}
}

// NewNode returns new node with value v
func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

// AppendToTail appends a new element to tail of a linked list
func (n *Node) AppendToTail(v interface{}) {
	for n.next != nil {
		n = n.next
	}
	n.next = &Node{Value: v}
}
