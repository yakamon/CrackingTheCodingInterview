package palindrome

// Node is a node of a linked list
type Node struct {
	Next  *Node
	Value interface{}
}

// Len returns the length of a linked list n
func (n *Node) Len() int {
	length := 1
	for n.Next != nil {
		length++
		n = n.Next
	}
	return length
}

// Values returns a slice of values of linked list n
func (n *Node) Values() []interface{} {
	var values []interface{}
	for n != nil {
		values = append(values, n.Value)
		n = n.Next
	}
	return values
}

// PushBack add a new node to the back of linked list n
func (n *Node) PushBack(v interface{}) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewByValue(v)
}

// PushBackSlice add new nodes to the back of linked list n
func (n *Node) PushBackSlice(l []interface{}) {
	for _, v := range l {
		n.PushBack(v)
	}
}

// NewByValue initializes a new node of a linked list with value v
func NewByValue(v interface{}) *Node {
	n := new(Node)
	n.Value = v
	return n
}

// NewBySlice initializes a new linked list with a slice of values l
func NewBySlice(l []interface{}) *Node {
	length := len(l)
	if length == 0 {
		return nil
	}

	n := NewByValue(l[0])
	if length == 1 {
		return n
	}

	for _, v := range l[1:] {
		n.PushBack(v)
	}
	return n
}

// Clone returns a clone of linked list n
func Clone(n *Node) *Node {
	return NewBySlice(n.Values())
}

// ReverseClone returns a reversed clone of linked list n
func ReverseClone(n *Node) *Node {
	var head *Node
	for n != nil {
		node := NewByValue(n.Value)
		node.Next = head
		head = node
		n = n.Next
	}
	return head
}

// IsEqual checks if two list l1 and l2 is equal or not
func IsEqual(l1, l2 *Node) bool {
	for l1 != nil && l2 != nil {
		if l1.Value != l2.Value {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
