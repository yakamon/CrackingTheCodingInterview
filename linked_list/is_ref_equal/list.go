package refequal

// Node is a node of linked list
type Node struct {
	next  *Node
	value interface{}
}

// Next returns the next node
func (n *Node) Next() *Node {
	return n.next
}

// Value returns the value of node
func (n *Node) Value() interface{} {
	return n.value
}

// Len returns the length of list
func (n *Node) Len() int {
	length := 1
	for n.Next() != nil {
		length++
		n = n.Next()
	}
	return length
}

// Tail returns the last node of list
func (n *Node) Tail() *Node {
	for n.Next() != nil {
		n = n.Next()
	}
	return n
}

// Values returns values of list
func (n *Node) Values() []interface{} {
	var vals []interface{}
	for n != nil {
		vals = append(vals, n.Value())
		n = n.Next()
	}
	return vals
}

// SetNext sets a node to the next to the current node
func (n *Node) SetNext(next *Node) {
	n.next = next
}

// PushBack add a new node to the tail of list
func (n *Node) PushBack(v interface{}) {
	for n.Next() != nil {
		n = n.Next()
	}
	n.SetNext(NewByValue(v))
}

// NewByValue creates a new list with value v
func NewByValue(v interface{}) *Node {
	return &Node{value: v}
}

// NewBySlice creates a new list with values l
func NewBySlice(l []interface{}) *Node {
	if len(l) == 0 {
		return nil
	}

	n := NewByValue(l[0])
	if len(l) == 1 {
		return n
	}

	for _, v := range l[1:] {
		n.PushBack(v)
	}
	return n
}
