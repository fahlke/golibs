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

	tests := []struct {
		name     string
		item     interface{}
		expected interface{}
	}{
		{"boolean", true, true},
		{"string", "foo", "foo"},
		{"signed integer", 0x1, 0x1},               //nolint:gonmd
		{"unsigned integer", uint(0x1), uint(0x1)}, //nolint:gonmd
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := Stack{}
			s.Push(tt.item)
			data, _ := s.Top()
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
		assert.Exactly(t, uint(0x0), s.Size()) //nolint:gomnd
	})

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		top, err := s.Top()
		assert.NoError(t, err)
		assert.Exactly(t, "foo", top)
		assert.Exactly(t, uint(0x1), s.Size()) //nolint:gomnd
	})

	t.Run("two items", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")
		top, err := s.Top()
		assert.NoError(t, err)
		assert.Exactly(t, "bar", top)
		assert.Exactly(t, uint(0x2), s.Size()) //nolint:gomnd
	})
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Stack{}

		_, err := s.Pop()
		assert.Error(t, ErrEmptyStack, err)
	})

	t.Run("one item", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")

		_, err := s.Pop()
		assert.NoError(t, err)
		assert.EqualValues(t, uint(0x0), s.Size()) //nolint:gomnd
	})

	t.Run("two items with one remaining", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")

		_, err := s.Pop()
		assert.NoError(t, err)
		assert.EqualValues(t, uint(0x1), s.Size()) //nolint:gomnd
	})

	t.Run("two items with empty stack", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		s.Push("foo")
		s.Push("bar")

		_, err := s.Pop()
		assert.NoError(t, err)
		assert.EqualValues(t, uint(0x1), s.Size()) //nolint:gomnd

		_, err = s.Pop()
		assert.NoError(t, err)
		assert.EqualValues(t, uint(0x0), s.Size()) //nolint:gomnd

		_, err = s.Pop()
		assert.Error(t, ErrEmptyStack, err)
		assert.EqualValues(t, uint(0x0), s.Size()) //nolint:gomnd
		assert.True(t, s.Empty())
	})
}
