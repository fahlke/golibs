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
		assert.NoError(t, s.Push("foo"))
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
		assert.NoError(t, s.Push("foo"))
		assert.Exactly(t, uint(0x1), s.Size()) //nolint:gomnd
	})

	t.Run("two items", func(t *testing.T) {
		t.Parallel()

		s := Stack{}
		assert.NoError(t, s.Push("foo"))
		assert.NoError(t, s.Push("bar"))
		assert.Exactly(t, uint(0x2), s.Size()) //nolint:gomnd
	})
}

func TestStack_Push(t *testing.T) { t.Parallel() }

func TestStack_Top(t *testing.T) { t.Parallel() }

func TestStack_Pop(t *testing.T) { t.Parallel() }
