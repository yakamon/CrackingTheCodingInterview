package rmcn

// Node is node of linked list
type Node struct {
	Next  *Node
	Value interface{}
}

// NewNode returns a new node of linked list
func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

// AppendToTail appends a new node to linked list with value v
func (n *Node) AppendToTail(v interface{}) {
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(v)
}

// AppendListToTail appends nodes to linked list
func (n *Node) AppendListToTail(list []interface{}) {
	for _, v := range list {
		n.AppendToTail(v)
	}
}

// Contains check if a list contains a value
func (n *Node) Contains(v interface{}) bool {
	if n == nil {
		return false
	}

	for n.Next != nil {
		if n.Value == v {
			return true
		}
		n = n.Next
	}

	return false
}
