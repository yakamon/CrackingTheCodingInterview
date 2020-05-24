package palindrome

// IsPalindrome checks if a linked list l is a palindrome or not
func IsPalindrome(l *Node) bool {
	len := l.Len() // O(N)
	half, isOdd := len/2, len%2 == 1
	halfValues := make([]interface{}, half, half)
	for i := half - 1; i >= 0; i-- { // O(N)
		halfValues[i] = l.Value
		l = l.Next
	}
	if isOdd {
		l = l.Next
	}
	for _, v := range halfValues { // O(N)
		if l.Value != v {
			return false
		}
		l = l.Next
	}
	return true
}

// IsPalindrome2 checks if the elements of a linked list l is a palindrome or not
func IsPalindrome2(l *Node) bool {
	reversedList := ReverseClone(l) // O(N)
	return IsEqual(l, reversedList) // O(N)
}

// IsPalindrome3 checks if the elements of a linked list l is a palindrome or not
func IsPalindrome3(l *Node) bool {
	slow, fast := l, l
	var values []interface{}
	for fast != nil && fast.Next != nil {
		values = append(values, slow.Value)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for i := len(values) - 1; i >= 0; i-- {
		if slow.Value != values[i] {
			return false
		}
		slow = slow.Next
	}
	return true
}

// IsPalindrome4 checks if the elements of a linked list l is a palindrome or not
func IsPalindrome4(l *Node) bool {
	length := l.Len()
	_, _, isPalindrome := isPalindrome4Recurse(l, length, false)
	return isPalindrome
}

func isPalindrome4Recurse(l *Node, length int, isPalindrome bool) (*Node, int, bool) {
	if l == nil || length <= 0 {
		return l, 0, true
	} else if length == 1 {
		return l.Next, 0, true
	}
	node, length, isPalindrome := isPalindrome4Recurse(l.Next, length-2, false)
	if !isPalindrome || node == nil {
		return node, 0, isPalindrome
	}
	return node.Next, 0, l.Value == node.Value
}
