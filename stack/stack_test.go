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

func TestStack_Size(t *testing.T) { t.Parallel() }

func TestStack_Push(t *testing.T) { t.Parallel() }

func TestStack_Top(t *testing.T) { t.Parallel() }

func TestStack_Pop(t *testing.T) { t.Parallel() }
