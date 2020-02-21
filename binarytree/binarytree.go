package binarytree

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/fahlke/golibs/util"
)

type binaryTree struct {
	mutex sync.RWMutex
	root  *binaryNode
}

type binaryNode struct {
	key   string
	value interface{}
	left  *binaryNode
	right *binaryNode
}

// Item is a key value pair of a node that will be returned by the iterator
type Item struct {
	Key   string
	Value interface{}
}

// New returns a binary tree.
func New() *binaryTree {
	bt := &binaryTree{}
	return bt
}

func (n *binaryNode) set(key string, value interface{}) {
	switch {
	case key < n.key:
		if n.left == nil {
			n.left = &binaryNode{key, value, nil, nil}
		} else {
			n.left.set(key, value)
		}
	case key > n.key:
		if n.right == nil {
			n.right = &binaryNode{key, value, nil, nil}
		} else {
			n.right.set(key, value)
		}
	default: // key == n.key
		n.value = value
	}
}

func (bt *binaryTree) Set(key string, value interface{}) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()

	if bt.root == nil {
		bt.root = &binaryNode{key: key, value: value}
	} else {
		bt.root.set(key, value)
	}
}

func (n *binaryNode) get(key string) (interface{}, error)  { return nil, nil }
func (bt *binaryTree) Get(key string) (interface{}, error) { return nil, nil }

func (n *binaryNode) delete(key string) (*binaryNode, error) { return nil, nil }
func (bt *binaryTree) Delete(key string) error               { return nil }

func (n *binaryNode) height() int {
	if n == nil {
		return 0
	}

	return 1 + util.Max(n.left.height(), n.right.height())
}

func (bt *binaryTree) Height() int {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()

	return bt.root.height()
}

func (n *binaryNode) iterate(ch chan<- Item) {}
func (bt *binaryTree) Iterate() <-chan Item  { ch := make(chan Item); return ch }

func treeView(buf io.Writer, node *binaryNode, level int, direction string) {
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

func (bt *binaryTree) String() string {
	bt.mutex.RLock()
	defer bt.mutex.RUnlock()

	buf := new(bytes.Buffer)

	treeView(buf, bt.root, 0, "root")

	return buf.String()
}
