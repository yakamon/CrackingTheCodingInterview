package refequal

// FindIntersection checks if two linked list l1 and l2 is equal by references
func FindIntersection(l1, l2 *Node) (*Node, bool) {
	if l1 == nil || l2 == nil {
		return nil, false
	}

	tail1, size1 := getTailAndSize(l1)
	tail2, size2 := getTailAndSize(l2)

	if tail1 != tail2 {
		return nil, false
	}

	if size1 < size2 {
		l2 = moveForwardByTimes(l2, size2-size1)
	} else {
		l1 = moveForwardByTimes(l1, size1-size2)
	}

	for l1 != l2 {
		l1 = l1.Next()
		l2 = l2.Next()
	}

	return l1, l1 != nil
}

func getTailAndSize(n *Node) (*Node, int) {
	size := 1
	for n.Next() != nil {
		size++
		n = n.Next()
	}
	return n, size
}

func moveForwardByTimes(n *Node, times int) *Node {
	for i := 0; i < times; i++ {
		n = n.Next()
	}
	return n
}
