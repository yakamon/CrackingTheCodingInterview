package numadd

// AddReverseListedNumbers adds two numbers represented by a linked list
func AddReverseListedNumbers(a, b *Node) *Node {
	var head, tail *Node
	carryDigit := 0
	for a != nil || b != nil {
		aval, bval := 0, 0
		if a != nil {
			aval = a.Value.(int)
		}
		if b != nil {
			bval = b.Value.(int)
		}
		sum := aval + bval + carryDigit
		currentDigit := sum % 10
		carryDigit = sum / 10
		currentNode := NewNode(currentDigit)

		if head == nil {
			head = currentNode
			tail = head
		} else {
			tail.Next = currentNode
			tail = currentNode
		}

		a, b = a.Next, b.Next
	}

	if carryDigit != 0 {
		tail.Next = NewNode(carryDigit)
	}

	return head
}

// AddReverseListedNumbers2 is the same as avobe
func AddReverseListedNumbers2(a, b *Node) *Node {
	return addReverseListedNumbers2(a, b, 0)
}

func addReverseListedNumbers2(a, b *Node, carry int) *Node {
	if a == nil && b == nil && carry == 0 {
		return nil
	}

	var anext, bnext *Node
	result := new(Node)
	value := carry
	if a != nil {
		value += a.Value.(int)
		anext = a.Next
	}
	if b != nil {
		value += b.Value.(int)
		bnext = b.Next
	}
	result.Value, carry = value%10, value/10

	if a != nil || b != nil {
		more := addReverseListedNumbers2(anext, bnext, carry)
		result.Next = more
	}
	return result
}
