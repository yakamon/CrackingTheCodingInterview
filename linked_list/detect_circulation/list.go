package circulation

// Node is node of linked list
type Node struct {
	next  *Node
	value interface{}
}

// Next returns the next node of node n
func (n *Node) Next() *Node {
	return n.next
}

// Value returns the value of node n
func (n *Node) Value() interface{} {
	return n.value
}

// Len returns the length of list from node n
func (n *Node) Len() int {
	length := 1
	for n.Next() != nil {
		n = n.Next()
		length++
	}
	return length
}

// Tail returns the last node of list n
func (n *Node) Tail() *Node {
	for n.Next() != nil {
		n  = n.Next()
	}
	return n
}

// SetNext sets the next node of the current node n
func (n *Node) SetNext(next *Node) {
	n.next = next
}

// PushBack add a new node initialized with value v to list
func (n *Node) PushBack(v interface{}) {
	for n.Next() != nil {
		n = n.Next()
	}
	n.SetNext(NewByValue(v))
}

// NewByValue returns a new list node initialized with value v
func NewByValue(v interface{}) *Node {
	return &Node{value: v}
}

// NewBySlice returns a new list with a list of values l
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
