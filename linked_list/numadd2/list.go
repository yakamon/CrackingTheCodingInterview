package numadd2

// Node is node of linked list
type Node struct {
	Next  *Node
	Value int
}

// NewNodeWithValue initialize a new linked list with a value
func NewNodeWithValue(v int) *Node {
	node := new(Node)
	node.Value = v
	return node
}

// NewListFromSlice initialize a new linked list with a slice
func NewListFromSlice(l []int) *Node {
	if len(l) == 0 {
		return nil
	}
	node := NewNodeWithValue(l[0])
	node.PushBackList(l[1:])
	return node
}

// Len returns the length of linked list
func (node *Node) Len() int {
	length := 1
	for node.Next != nil {
		length++
		node = node.Next
	}
	return length
}

// Values returns a slice of values of linked list
func (node *Node) Values() []int {
	var values []int
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

// PushBack add a new node to the back of linked list
func (node *Node) PushBack(v int) *Node {
	for node.Next != nil {
		node = node.Next
	}
	node.Next = NewNodeWithValue(v)
	return node
}

// PushBackList add new nodes from a list to the back of linked list
func (node *Node) PushBackList(l []int) *Node {
	for _, v := range l {
		node.PushBack(v)
	}
	return node
}

// PushFront add a new node to the front of linked list
func PushFront(node *Node, v int) *Node {
	front := NewNodeWithValue(v)
	front.Next = node
	return front
}
