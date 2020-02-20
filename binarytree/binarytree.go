package binarytree

import (
	"sync"
)

// BinaryTree represents a binary search tree implementation.
type BinaryTree struct {
	mutex sync.RWMutex
	root  *node
}

// Item is a key value pair of a node that will be returned by the iterator
type Item struct {
	Key   string
	Value interface{}
}

type node struct {
	key   string
	value interface{}
	left  *node
	right *node
}

func (n *node) set(key string, value interface{})        {}
func (bt *BinaryTree) Set(key string, value interface{}) {}

func (n *node) get(key string) (interface{}, error)        { return nil, nil }
func (bt *BinaryTree) Get(key string) (interface{}, error) { return nil, nil }

func (n *node) delete(key string) (*node, error) { return nil, nil }
func (bt *BinaryTree) Delete(key string) error   { return nil }

func (n *node) height() int        { return 0 }
func (bt *BinaryTree) Height() int { return 0 }

func (n *node) iterate(ch chan<- Item)      {}
func (bt *BinaryTree) Iterate() <-chan Item { ch := make(chan Item); return ch }
