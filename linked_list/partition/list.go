package partition

// Node is node of linked list
type Node struct {
	Next  *Node
	Value interface{}
}

// NewNode returns a new node of list
func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

// AppendToTail append a node to list
func (n *Node) AppendToTail(v interface{}) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(v)
}

// AppendListToTail append a list to list
func (n *Node) AppendListToTail(list []interface{}) {
	for _, v := range list {
		n.AppendToTail(v)
	}
}

// Values returns all values of linked list as a slice of interface{}
func (n *Node) Values() []interface{} {
	var values []interface{}
	for n != nil {
		values = append(values, n.Value)
		n = n.Next
	}
	return values
}
