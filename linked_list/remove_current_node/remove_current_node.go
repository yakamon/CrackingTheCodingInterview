package rmcn

// RemoveCurrentNode remove current node from the linked list and returns true if success
func RemoveCurrentNode(n *Node) bool {
	if n == nil || n.Next == nil {
		return false
	}
	next := n.Next
	n.Value = next.Value
	n.Next = next.Next
	return true
}
