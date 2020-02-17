package linkedlist_test

import (
	"fmt"

	"github.com/fahlke/golibs/linkedlist"
)

func ExampleLinkedList_Append() {
	list := linkedlist.LinkedList{}

	list.Append("foo")
	list.Append("bar")
	list.Append("baz")
	// Output:
}

func ExampleLinkedList_Iterate() {
	list := linkedlist.LinkedList{}

	list.Append("foo")
	list.Append("bar")
	list.Append("baz")

	for node := range list.Iterate() {
		fmt.Println(node)
	}
	// Output:
	// foo
	// bar
	// baz
}

func ExampleLinkedList_GetNth() {
	list := linkedlist.LinkedList{}

	list.Append("foo")
	list.Append("bar")
	list.Append("baz")

	node, _ := list.GetNth(1)
	fmt.Println(node)
	// Output:
	// bar
}

func ExampleLinkedList_RemoveBeginning() {
	list := linkedlist.LinkedList{}

	list.Append("foo")
	list.Append("bar")
	list.Append("baz")

	list.RemoveBeginning()

	for node := range list.Iterate() {
		fmt.Println(node)
	}
	// Output:
	// bar
	// baz
}
