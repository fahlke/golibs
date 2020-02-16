package queue_test

import (
	"fmt"

	"github.com/fahlke/golibs/queue"
)

func ExampleQueue_Empty() {
	q := queue.Queue{}

	fmt.Println(q.Empty())
	// Output:
	// true
}

func ExampleQueue_Size() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")

	fmt.Println(q.Size())
	// Output:
	// 2
}

func ExampleQueue_Push() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")
}

func ExampleQueue_Pop() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")

	item, _ := q.Pop()
	fmt.Println(item)
	// Output:
	// foo
}

func ExampleQueue_Front() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")

	front, _ := q.Front()
	fmt.Println(front)
	// Output:
	// foo
}

func ExampleQueue_Back() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")

	back, _ := q.Back()
	fmt.Println(back)
	// Output:
	// baz
}
