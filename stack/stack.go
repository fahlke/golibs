package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	sync.Mutex
	items []interface{}
}

// ErrEmptyStack returns on an invalid operation on an empty stack
var ErrEmptyStack = errors.New("empty stack")

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() uint {
	return uint(len(s.items))
}

func (s *Stack) Push(item interface{}) error {
	s.items = append(s.items, item)

	return nil
}

func (s *Stack) Top() (interface{}, error) { return nil, nil }

func (s *Stack) Pop() (interface{}, error) { return nil, nil }
