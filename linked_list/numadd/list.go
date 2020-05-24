package numadd

// Node is a node of linked list
type Node struct {
	Next  *Node
	Value interface{}
}

// NewNode initializes a new node with value v
func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

// NewListFromSlice initializes a new linked list with values of slice
func NewListFromSlice(list []interface{}) *Node {
	node := NewNode(list[0])
	node.AppendListToTail(list[1:])
	return node
}

// AppendToTail append a node to linked list
func (n *Node) AppendToTail(v interface{}) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(v)
}

// AppendListToTail appends list values of list to linked list
func (n *Node) AppendListToTail(list []interface{}) {
	for _, v := range list {
		n.AppendToTail(v)
	}
}

// Values returns all values of linked list as a slice
func (n *Node) Values() []interface{} {
	var values []interface{}
	for n != nil {
		values = append(values, n.Value)
		n = n.Next
	}
	return values
}
