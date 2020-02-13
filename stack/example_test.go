package stack_test

import (
	"fmt"

	"github.com/fahlke/golibs/stack"
)

func ExampleStack_Empty() {
	s := stack.Stack{}
	fmt.Println(s.Empty())
	// Output:
	// true
}

func ExampleStack_Size() {
	s := stack.Stack{}
	fmt.Println(s.Size())
	// Output:
	// 0
}

func ExampleStack_Push() {
	s := stack.Stack{}

	s.Push("foo")
	s.Push("bar")
	// Output:
}

func ExampleStack_Top() {
	s := stack.Stack{}
	s.Push("foo")
	s.Push("bar")

	item, _ := s.Top()
	fmt.Println(item)
	// Output:
	// bar
}

func ExampleStack_Pop() {
	s := stack.Stack{}

	s.Push("foo")
	s.Push("bar")

	for !s.Empty() {
		item, _ := s.Pop()
		fmt.Println(item)
	}
	// Output:
	// bar
	// foo
}
