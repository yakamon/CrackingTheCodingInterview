package partition

// ByPivot1 sorts linked list based on pivot x
func ByPivot1(n *Node, x int) *Node {
	var beforeStart, beforeEnd, afterStart, afterEnd *Node
	for n != nil {
		next := n.Next
		n.Next = nil
		if n.Value.(int) < x {
			if beforeStart == nil {
				beforeStart = n
				beforeEnd = beforeStart
			} else {
				beforeEnd.Next = n
				beforeEnd = n
			}
		} else {
			if afterStart == nil {
				afterStart = n
				afterEnd = afterStart
			} else {
				afterEnd.Next = n
				afterEnd = n
			}
		}
		n = next
	}

	if beforeStart == nil {
		return afterStart
	}

	beforeEnd.Next = afterStart
	return beforeStart
}

// ByPivot2 sorts linked list based on pibot x
func ByPivot2(n *Node, x int) *Node {
	head, tail := n, n
	for n != nil {
		next := n.Next
		n.Next = nil
		if n.Value.(int) < x {
			n.Next = head
			head = n
		} else {
			tail.Next = n
			tail = n
		}
		n = next
	}

	return head
}
