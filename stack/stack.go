package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	sync.RWMutex
	items []interface{}
}

// ErrEmptyStack returns on an invalid operation on an empty stack
var ErrEmptyStack = errors.New("empty stack")

func (s *Stack) Empty() bool {
	s.RLock()
	defer s.RUnlock()

	return len(s.items) == 0
}

func (s *Stack) Size() uint {
	s.RLock()
	defer s.RUnlock()

	return uint(len(s.items))
}

func (s *Stack) Push(item interface{}) error {
	s.Lock()
	defer s.Unlock()

	s.items = append(s.items, item)

	return nil
}

func (s *Stack) Top() (interface{}, error) {
	s.RLock()
	defer s.RUnlock()

	if s.Empty() {
		return nil, ErrEmptyStack
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	s.Lock()
	defer s.Unlock()

	return nil, nil
}
