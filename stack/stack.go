package stack

import (
	"errors"
	"sync"
)

// Stack represents a stack implementation
type Stack struct {
	mutex sync.RWMutex
	items []interface{}
}

// ErrEmptyStack will be returned on an invalid operation on an empty stack.
var ErrEmptyStack = errors.New("empty stack")

// Empty returns true if the underlying stack is empty, false otherwise.
func (s *Stack) Empty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.items) == 0
}

// Size returns the number of elements in the underlying stack.
func (s *Stack) Size() uint {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return uint(len(s.items))
}

// Push is putting the given element value to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items = append(s.items, item)
}

// Top returns the top element in the stack without removing it. That is the most recently pushed element.
func (s *Stack) Top() (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if len(s.items) == 0 {
		return nil, ErrEmptyStack
	}

	return s.items[len(s.items)-1], nil
}

// Pop removes the top element from the stack and returns it. It returns an error if the stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == 0 {
		return nil, ErrEmptyStack
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil
}
