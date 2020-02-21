package binarytree

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/fahlke/golibs/util"
)

// BinaryTree represents a binary search tree implementation.
type BinaryTree struct {
	mutex sync.RWMutex
	root  *node
}

type node struct {
	key   string
	value interface{}
	left  *node
	right *node
}

// Item is a key value pair of a node that will be returned by the iterator
type Item struct {
	Key   string
	Value interface{}
}

func New() *BinaryTree {
	bt := &BinaryTree{}
	return bt
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

func (n *node) height() int {
	if n == nil {
		return 0
	}

	return 1 + util.Max(n.left.height(), n.right.height())
}

func (bt *BinaryTree) Height() int {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()

	return bt.root.height()
}

func (n *node) iterate(ch chan<- Item)      {}
func (bt *BinaryTree) Iterate() <-chan Item { ch := make(chan Item); return ch }

func treeView(buf io.Writer, node *node, level int, direction string) {
	if node == nil {
		return
	}

	for i := 0; i < level; i++ {
		_, _ = fmt.Fprint(buf, " ")
	}

	_, _ = fmt.Fprintf(buf, "%s: %s\n", direction, node.key)

	treeView(buf, node.left, 2+level, "left")   //nolint:gomnd
	treeView(buf, node.right, 2+level, "right") //nolint:gomnd
}

func (bt *BinaryTree) String() string {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()

	buf := new(bytes.Buffer)

	treeView(buf, bt.root, 0, "root")

	return buf.String()
}
