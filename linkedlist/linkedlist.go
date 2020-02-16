package linkedlist

import "sync"

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
func (l *LinkedList) Size() uint { return 0 }

// Empty ...
func (l *LinkedList) Empty() bool { return true }

// Append ...
func (l *LinkedList) Append(data interface{}) {}

// Remove ...
func (l *LinkedList) Remove(data interface{}) error { return nil }

// GetNth ...
func (l *LinkedList) GetNth(idx int) (interface{}, error) { return nil, nil }

// Swap ...
func (l *LinkedList) Swap(idx1, idx2 int) error { return nil }

// Iterate ...
func (l *LinkedList) Iterate() <-chan interface{} { return nil }

// InsertBeginning ...
func (l *LinkedList) InsertBeginning(data interface{}) {}

// InsertAfter ...
func (l *LinkedList) InsertAfter(data interface{}, idx int) error { return nil }

// RemoveBeginning ...
func (l *LinkedList) RemoveBeginning() error { return nil }

// RemoveAfter ...
func (l *LinkedList) RemoveAfter(idx int) error { return nil }
