package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Empty(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		assert.Equal(t, true, s.Empty())
	})

	t.Run("not empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		assert.Equal(t, false, s.Empty())
	})
}

func TestStack_Size(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		assert.Exactly(t, uint(0x0), s.Size()) //nolint:gomnd
	})

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		assert.Exactly(t, uint(0x1), s.Size()) //nolint:gomnd
	})

	t.Run("two items", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")
		assert.Exactly(t, uint(0x2), s.Size()) //nolint:gomnd
	})
}

func TestStack_Push(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		foo string
		bar bool
	}

	tests := []struct {
		name     string
		item     interface{}
		expected interface{}
	}{
		{"a boolean", true, true},
		{"a string", "foo", "foo"},
		{"a signed integer", 0x1, 0x1},               //nolint:gomnd
		{"a unsigned integer", uint(0x1), uint(0x1)}, //nolint:gomnd
		{"a custom struct", testStruct{"baz", false}, testStruct{}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := Stack{}
			s.Push(tt.item)
			data, err := s.Top()
			assert.NoError(t, err)
			assert.False(t, s.Empty())
			assert.IsType(t, tt.expected, data)
		})
	}
}

func TestStack_Top(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		_, err := s.Top()
		assert.Error(t, err, "empty stack")
	})

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		top, err := s.Top()
		assert.NoError(t, err)
		assert.Exactly(t, "foo", top)
	})

	t.Run("two items", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")
		top, err := s.Top()
		assert.NoError(t, err)
		assert.Exactly(t, "bar", top)
	})
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}

		item, err := s.Pop()
		assert.Nil(t, item)
		assert.Error(t, ErrEmptyStack, err)
	})

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")

		item, _ := s.Pop()
		assert.Equal(t, "foo", item)
	})

	t.Run("two items", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")

		item, _ := s.Pop()
		assert.Equal(t, "bar", item)

		item, _ = s.Pop()
		assert.Equal(t, "foo", item)
	})
}
