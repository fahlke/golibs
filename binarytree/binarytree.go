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

func (n *node) set(key string, value interface{}) {
	switch {
	case key < n.key:
		if n.left == nil {
			n.left = &node{key, value, nil, nil}
		} else {
			n.left.set(key, value)
		}
	case key > n.key:
		if n.right == nil {
			n.right = &node{key, value, nil, nil}
		} else {
			n.right.set(key, value)
		}
	default: // key == n.key
		n.value = value
	}
}

func (bt *BinaryTree) Set(key string, value interface{}) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()

	if bt.root == nil {
		bt.root = &node{key: key, value: value}
	} else {
		bt.root.set(key, value)
	}
}

func (n *node) get(key string) (interface{}, error)        { return nil, nil }
func (bt *BinaryTree) Get(key string) (interface{}, error) { return nil, nil }

func (n *node) delete(key string) (*node, error) { return nil, nil }
func (bt *BinaryTree) Delete(key string) error   { return nil }

func (n *node) height() int        { return 0 }
func (bt *BinaryTree) Height() int { return 0 }

func (n *node) iterate(ch chan<- Item)      {}
func (bt *BinaryTree) Iterate() <-chan Item { ch := make(chan Item); return ch }
