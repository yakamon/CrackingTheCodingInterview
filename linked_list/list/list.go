package list

// Element is an element of a linked list
type Element struct {
	prev, next *Element
	list       *List
	Value      interface{}
}

// Next returns the next list element or nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List is a linked list
type List struct {
	root Element
	len  int
}

// Init initializes or clears list l
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list
func New() *List { return new(List).Init() }

// Len returns the number of elements of list len (Complexity: O(1))
func (l *List) Len() int { return l.len }

// Front returns the first element of list l or nil of the list is empty
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil of the list is empty
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e
func (l *List) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a wrapper for insert
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

// move moves e to next to at and returns e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

// Remove removes e from l if e is an element of list l
// It returns the element value e.Value
// The element must not be nil
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e
// If mark is not an element of l, the list is not modified
// The mark must not be nil
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e
// If mark is not an element of l, the list is not modified
// The mark mut not be nil
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l
// If e is not an element of list l, the list is not modified
// The element must not be nil
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l
// If e is not an element of list l, the list is not modified
// The element must not be nil
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark
// If e or mark is not an element of l, or e == mark, the list is not modified
// The element must not be nil
func (l *List) MoveBefore(e, mark *Element) {
	if e.list != l || mark.list != nil || e == mark {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new positon after mark
// If e or mark is not an element of l, or e == mark, the list is not modified
// The element must not be nil
func (l *List) MoveAfter(e, mark *Element) {
	if e.list != l || mark.list != nil || e == mark {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of an other list at the back of list l
// The list l and other may be the same
// They must not be nil
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	// 一番前から順番に後方に insert していく
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l
// The list l and other may be the same
// They must not be nil
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	// 一番後ろから順番に前方に insert していく
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}