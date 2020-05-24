package numadd2

// AddLists add two numbers represented by linked list
func AddLists(l1, l2 *Node) *Node {
	len1, len2 := l1.Len(), l2.Len()
	if len1 < len2 {
		l1 = padList(l1, len2-len1)
	} else if len1 > len2 {
		l2 = padList(l2, len1-len2)
	}

	sum := addLists(l1, l2)
	if sum.Carry == 0 {
		return sum.Node
	}

	return PushFront(sum.Node, sum.Carry)
}

type partialSum struct {
	Node  *Node
	Carry int
}

func addLists(l1, l2 *Node) *partialSum {
	if l1 == nil && l2 == nil {
		sum := new(partialSum)
		return sum
	}
	sum := addLists(l1.Next, l2.Next)
	value := l1.Value + l2.Value + sum.Carry
	sum.Node = PushFront(sum.Node, value%10)
	sum.Carry = value / 10
	return sum
}

func padList(l *Node, len int) *Node {
	for i := 0; i < len; i++ {
		l = PushFront(l, 0)
	}
	return l
}
