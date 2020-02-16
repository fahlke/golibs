package linkedlist

import (
	"errors"
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

	t.Run("remove non-existing item", func(t *testing.T) {
		l := LinkedList{}
		l.Append("foo")
		l.Append("bar")

		err := l.Remove("baz")
		assert.Error(t, err)
		assert.NotEqual(t, nil, l.head.next)
		assert.EqualValues(t, 0x2, l.size)
		assert.EqualValues(t, "foo", l.head.data)
		assert.EqualValues(t, "bar", l.head.next.data)
	})

	t.Run("remove first item", func(t *testing.T) {
		l := LinkedList{}
		l.Append("foo")
		l.Append("bar")
		l.Append("baz")

		err := l.Remove("foo")
		assert.NoError(t, err)
		assert.NotEqual(t, nil, l.head.next)
		assert.EqualValues(t, 0x2, l.size)
		assert.EqualValues(t, "bar", l.head.data)
		assert.EqualValues(t, "baz", l.head.next.data)
	})

	t.Run("remove middle item", func(t *testing.T) {
		l := LinkedList{}
		l.Append("foo")
		l.Append("bar")
		l.Append("baz")

		err := l.Remove("bar")
		assert.NoError(t, err)
		assert.NotEqual(t, nil, l.head.next)
		assert.EqualValues(t, 0x2, l.size)
		assert.EqualValues(t, "foo", l.head.data)
		assert.EqualValues(t, "baz", l.head.next.data)
	})

	t.Run("remove last item", func(t *testing.T) {
		l := LinkedList{}
		l.Append("foo")
		l.Append("bar")
		l.Append("baz")

		err := l.Remove("baz")
		assert.NoError(t, err)
		assert.NotEqual(t, nil, l.head.next)
		assert.EqualValues(t, 0x2, l.size)
		assert.EqualValues(t, "foo", l.head.data)
		assert.EqualValues(t, "bar", l.head.next.data)
	})
}

func TestLinkedList_GetNth(t *testing.T) {
	t.Parallel()

	l := LinkedList{}
	l.Append("foo")
	l.Append("bar")
	l.Append("baz")
	l.Append("...")
	l.Append("end")

	t.Run("out of bounds", func(t *testing.T) {
		item, err := l.GetNth(0xf) //nolint:gomnd
		assert.Error(t, errors.New("access out of bounds"), err)
		assert.EqualValues(t, nil, item)
	})

	t.Run("1st item", func(t *testing.T) {
		item, err := l.GetNth(0) //nolint:gomnd
		assert.NoError(t, err)
		assert.EqualValues(t, "foo", item)
	})

	t.Run("3rd item", func(t *testing.T) {
		item, err := l.GetNth(2) //nolint:gomnd
		assert.NoError(t, err)
		assert.EqualValues(t, "baz", item)
	})

	t.Run("5th item", func(t *testing.T) {
		item, err := l.GetNth(4) //nolint:gomnd
		assert.NoError(t, err)
		assert.EqualValues(t, "end", item)
	})
}

func TestLinkedList_Swap(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_Iterate(t *testing.T) {
	t.Parallel()

	l := LinkedList{}

	l.Append("foo")
	l.Append("bar")
	l.Append("baz")

	var idx int

	for node := range l.Iterate() {
		switch idx {
		case 0:
			assert.Equal(t, "foo", node)
		case 1: //nolint:gomnd
			assert.Equal(t, "bar", node)
		case 2: //nolint:gomnd
			assert.Equal(t, "baz", node)
		}

		idx++
	}
}

func TestLinkedList_InsertBeginning(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestLinkedList_InsertAfter(t *testing.T) {
	t.Parallel()

	l := LinkedList{}
	l.Append("foo")
	l.Append("bar")
	l.Append("baz")

	t.Run("out of bounds", func(t *testing.T) {
		err := l.InsertAfter("fail", 0x3) //nolint:gomnd
		assert.Error(t, errors.New("access out of bounds"), err)
		assert.EqualValues(t, uint(0x3), l.size) //nolint:gomnd
	})

	t.Run("after 1st item", func(t *testing.T) {
		err := l.InsertAfter("after first", 0x0) //nolint:gomnd
		assert.NoError(t, err)
		assert.EqualValues(t, 0x4, l.Size())
		assert.EqualValues(t, "after first", l.head.next.data)
	})

	t.Run("after last item", func(t *testing.T) {
		err := l.InsertAfter("after last", 0x3) //nolint:gomnd
		assert.NoError(t, err)
		assert.EqualValues(t, 0x5, l.Size()) //nolint:gomnd
		assert.EqualValues(t, "after last", l.head.next.next.next.next.data)
	})
}

func TestLinkedList_RemoveBeginning(t *testing.T) {
	t.Parallel()

	l := LinkedList{}
	l.Append("foo")
	l.Append("bar")
	l.Append("baz")

	t.Run("two remaining", func(t *testing.T) {
		l.RemoveBeginning()
		assert.EqualValues(t, uint(0x2), l.size) //nolint:gomnd
		assert.EqualValues(t, "bar", l.head.data)
		assert.EqualValues(t, "baz", l.head.next.data)
	})

	t.Run("one remaining", func(t *testing.T) {
		l.RemoveBeginning()
		assert.EqualValues(t, uint(0x1), l.size) //nolint:gomnd
		assert.EqualValues(t, "baz", l.head.data)
	})

	t.Run("empty", func(t *testing.T) {
		l.RemoveBeginning()
		assert.EqualValues(t, uint(0x0), l.size) //nolint:gomnd
	})
}

func TestLinkedList_RemoveAfter(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}
