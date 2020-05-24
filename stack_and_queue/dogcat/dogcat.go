package dogcat

import (
	"errors"
	"time"
)

type animal struct {
	registeredAt int64
}

// Animal is interface for animal
type Animal interface {
	RegisteredAt() int64
}

// Dog is dog.
type Dog animal

// NewDog returns a new dog instance.
func NewDog() *Dog {
	return &Dog{time.Now().UnixNano()}
}

// RegisteredAt returns Dog.registered.
func (d *Dog) RegisteredAt() int64 {
	return d.registeredAt
}

// Cat is cat.
type Cat animal

// NewCat returns a new cat instance.
func NewCat() *Cat {
	return &Cat{time.Now().UnixNano()}
}

// RegisteredAt returns Cat.registeredAt.
func (c *Cat) RegisteredAt() int64 {
	return c.registeredAt
}

// DogCat is an extended queue structure.
type DogCat struct {
	dogs *queue
	cats *queue
}

// New returns a new DogCat instance.
func New() *DogCat {
	return &DogCat{newQueue(), newQueue()}
}

// Enqueue adds a new Animal to the back of DogCat.
func (q *DogCat) Enqueue(a Animal) error {
	switch a.(type) {
	case *Dog:
		q.dogs.Enqueue(a)
	case *Cat:
		q.cats.Enqueue(a)
	default:
		return errors.New("Unexpected type")
	}
	return nil
}

// DequeueAny retrieves the oldest animal from DogCat
func (q *DogCat) DequeueAny() (Animal, bool) {
	dog, okDog := q.dogs.Head()
	cat, okCat := q.cats.Head()
	if !okDog && !okCat {
		return nil, false
	}
 
	if !okCat || dog.(*Dog).RegisteredAt() < cat.(*Cat).RegisteredAt() {
		dog, ok := q.dogs.Dequeue()
		return dog.(*Dog), ok
	}
	cat, ok := q.cats.Dequeue()
	return cat.(*Cat), ok
}

// DequeueDog retrieves the oldest dog from DogCat
func (q *DogCat) DequeueDog() (Animal, bool) {
	dog, ok := q.dogs.Dequeue()
	return dog.(*Dog), ok
}

// DequeueCat retrieves the oldest cat from DogCat
func (q *DogCat) DequeueCat() (Animal, bool) {
	cat, ok := q.cats.Dequeue()
	return cat.(*Cat), ok
}

type queue struct {
	head *queueNode
	tail *queueNode
	size int
}

type queueNode struct {
	value interface{}
	back  *queueNode
}

func newQueue() *queue {
	return &queue{nil, nil, 0}
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func (q *queue) Enqueue(value interface{}) {
	n := &queueNode{value, nil}
	if q.tail != nil {
		q.tail.back = n
	}
	q.tail = n
	if q.head == nil {
		q.head = q.tail
	}
	q.size++
}

func (q *queue) Dequeue() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}

	n := q.head
	q.head = n.back
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return n.value, true
}

// Peek returns the value of first element of queue.
func (q *queue) Head() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}

	return q.head.value, true
}
