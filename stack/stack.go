package stack

import (
	"sync"
)

type Stack struct {
	sync.Mutex
	items []interface{}
}

func (s *Stack) Empty() bool { return true }

func (s *Stack) Size() int { return 0 }

func (s *Stack) Push(item interface{}) error { return nil }

func (s *Stack) Top() (interface{}, error) { return nil, nil }

func (s *Stack) Pop() (interface{}, error) { return nil, nil }
