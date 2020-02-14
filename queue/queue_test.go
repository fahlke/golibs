package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Empty(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		q := Queue{}
		assert.True(t, q.Empty())
	})

	t.Run("non-empty", func(t *testing.T) {
		t.Parallel()

		q := Queue{}
		q.Push("foo")
		assert.False(t, q.Empty())
	})
}

func TestQueue_Size(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestQueue_Push(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestQueue_Pop(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestQueue_Front(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}

func TestQueue_Back(t *testing.T) {
	t.Parallel()
	t.Error("implement me")
}
