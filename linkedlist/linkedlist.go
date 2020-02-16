package linkedlist

import (
	"errors"
	"sync"
)

var (
	// ErrNotFound means that the requested item does not exist
	ErrNotFound = errors.New("item not found")
	// ErrOutOfBounds means that an access outside of the list was attempted
	ErrOutOfBounds = errors.New("access out of bounds")
)

// LinkedList ...
type LinkedList struct {
	mutex sync.RWMutex
	head  *node
	size  uint
}

type node struct {
	next *node
	data interface{}
}

// Size ...
func (l *LinkedList) Size() uint {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	return l.size
}

// Empty ...
func (l *LinkedList) Empty() bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	return l.size == 0
}

// Append ...
func (l *LinkedList) Append(data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newNode := &node{
		data: data,
	}

	if l.head == nil {
		l.head = newNode
		l.size++

		return
	}

	var currentNode *node
	for currentNode = l.head; currentNode.next != nil; currentNode = currentNode.next {
	}

	currentNode.next = newNode
	l.size++
}

// Remove ...
func (l *LinkedList) Remove(data interface{}) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var previousNode *node

	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.data == data {
			if previousNode == nil {
				l.head = currentNode.next
			} else {
				previousNode.next = currentNode.next
			}

			l.size--

			return nil
		}

		previousNode = currentNode
	}

	return ErrNotFound
}

// GetNth ...
func (l *LinkedList) GetNth(idx uint) (interface{}, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if idx > l.size {
		return nil, ErrOutOfBounds
	}

	currentNode := l.head
	for i := uint(0x0); i < idx; i++ { //nolint:gomnd
		if currentNode.next != nil {
			currentNode = currentNode.next
		}
	}

	return currentNode.data, nil
}

// Swap ...
func (l *LinkedList) Swap(idx1, idx2 int) error { return nil }

// Iterate ...
func (l *LinkedList) Iterate() <-chan interface{} {
	ch := make(chan interface{}, l.size)
	l.mutex.RLock()

	go func() {
		for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
			ch <- currentNode.data
		}
		l.mutex.RUnlock()
		close(ch)
	}()

	return ch
}

// InsertBeginning ...
func (l *LinkedList) InsertBeginning(data interface{}) {}

// InsertAfter ...
func (l *LinkedList) InsertAfter(data interface{}, idx uint) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if idx >= l.size {
		return ErrOutOfBounds
	}

	newNode := &node{
		data: data,
	}

	currentNode := l.head

	var i = uint(0x0) //nolint:gomnd

	for currentNode.next != nil {
		nextNode := currentNode.next

		if i == idx {
			currentNode.next = newNode
			newNode.next = nextNode

			l.size++

			return nil
		}

		currentNode = nextNode
		i++
	}

	currentNode.next = newNode
	l.size++

	return nil
}

// RemoveBeginning ...
func (l *LinkedList) RemoveBeginning() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.head.next != nil {
		l.head = l.head.next
	} else {
		l.head = nil
	}

	l.size--
}

// RemoveAfter ...
func (l *LinkedList) RemoveAfter(idx int) error { return nil }
