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

func ExampleQueue_Push_Pop() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")

	item, _ := q.Pop()
	fmt.Println(item)
	// Output:
	// foo
}

func ExampleQueue_Push_Front_Back() {
	q := queue.Queue{}
	q.Push("foo")
	q.Push("bar")
	q.Push("baz")

	front, _ := q.Front()
	back, _ := q.Back()
	fmt.Println(front)
	fmt.Println(back)
	// Output:
	// foo
	// baz
}
