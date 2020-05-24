package mt

// Node is a node of BST.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func newNode(value int) *Node {
	return &Node{value, nil, nil}
}

// CreateBSTFromSortedSlice creates binary search tree from sorted int slice.
func CreateBSTFromSortedSlice(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		return newNode(s[0])
	}

	mid := len(s) / 2
	n := newNode(s[mid])
	n.Left = CreateBSTFromSortedSlice(s[:mid-1])
	n.Right = CreateBSTFromSortedSlice(s[mid:])
	return n
}
