package circulation

// DetectCirculation detects a circulation of linked list
// and returns the first node of circulating list and if the l is circulating or not
func DetectCirculation(l *Node) (*Node, bool) {
	// 1. save nodes to map
	nodes := map[*Node]bool{}
	for l != nil {
		if _, found := nodes[l]; found {
			return l, found
		}
		nodes[l] = true
		l = l.Next()
	}
	return nil, false
}

// DetectCirculation2 detects a circulation of linked list
func DetectCirculation2(l *Node) (*Node, bool) {
	slow, fast := l, l
	for fast != nil && fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next().Next()
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next() == nil {
		return nil, false
	}

	slow = l
	for slow != fast {
		slow = slow.Next()
		fast = fast.Next()
	}

	return slow, true
}
