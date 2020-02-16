package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Size(t *testing.T) {
	t.Parallel()

	l := LinkedList{}

	t.Run("empty", func(t *testing.T) {
		assert.EqualValues(t, 0, l.Size())
	})

	t.Run("one item", func(t *testing.T) {
		l.Append("foo")
		assert.EqualValues(t, 1, l.Size())
	})

	t.Run("two items", func(t *testing.T) {
		l.Append("baz")
		assert.EqualValues(t, 0x2, l.Size())
	})
}

func TestLinkedList_Empty(t *testing.T) {
	t.Parallel()

	l := LinkedList{}

	t.Run("empty", func(t *testing.T) {
		assert.EqualValues(t, true, l.Empty())
	})

	t.Run("not empty", func(t *testing.T) {
		l.Append("foo")
		assert.EqualValues(t, false, l.Empty())
	})
}

func TestLinkedList_Append(t *testing.T) {
	t.Parallel()

	l := LinkedList{}

	t.Run("one item", func(t *testing.T) {
		l.Append("foo")
		assert.NotEqual(t, nil, l.head)
		assert.EqualValues(t, "foo", l.head.data)
	})

	t.Run("two items", func(t *testing.T) {
		l.Append("bar")
		assert.NotEqual(t, nil, l.head.next)
		assert.EqualValues(t, "bar", l.head.next.data)
	})

	t.Run("three items", func(t *testing.T) {
		l.Append("baz")
		assert.NotEqual(t, nil, l.head.next.next)
		assert.EqualValues(t, "baz", l.head.next.next.data)
	})
}

func TestLinkedList_Remove(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_GetNth(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_Swap(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_Iterate(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_InsertBeginning(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_InsertAfter(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_RemoveBeginning(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_RemoveAfter(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}
