package rmdup

import "container/list"

// RemoveDuplicateElements removes duplicate elements in linked list l
func RemoveDuplicateElements(l *list.List) {
	dupMap := map[interface{}]bool{}
	for e := l.Front(); e != nil; {
		if _, exists := dupMap[e.Value]; !exists {
			dupMap[e.Value] = true
			e = e.Next()
		} else {
			n := e.Next()
			l.Remove(e)
			e = n
		}
	}
}

// RemoveDuplicateElements2 removes duplicate elements in linked list l
func RemoveDuplicateElements2(l *list.List) {
	for i, origin := 0, l.Front(); i < l.Len()-1; i, origin = i+1, origin.Next() {
		e := origin.Next()
		for e != nil {
			if e.Value == origin.Value {
				n := e.Next()
				l.Remove(e)
				e = n
			} else {
				e = e.Next()
			}
		}
	}
}
