package queue

import (
	"errors"
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

	golden := []struct {
		name  string
		items []string
		size  int
	}{
		{"empty", []string{}, 0x0},
		{"one item", []string{"foo"}, 0x1},
		{"two items", []string{"foo", "bar"}, 0x1},
	}

	for _, tt := range golden {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := Queue{}
			for _, item := range tt.items {
				q.Push(item)
			}
			assert.EqualValues(t, tt.size, q.Size())
		})
	}
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

	type expected struct {
		item interface{}
		err  error
	}

	golden := []struct {
		name     string
		items    []string
		expected expected
	}{
		{"empty", []string{}, expected{nil, errors.New("empty queue")}},
		{"one item", []string{"foo"}, expected{"foo", nil}},
		{"two items", []string{"foo", "bar"}, expected{"bar", nil}},
		{"three items", []string{"foo", "bar", "baz"}, expected{"baz", nil}},
	}

	for _, tt := range golden {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := Queue{}

			for _, item := range tt.items {
				q.Push(item)
			}

			item, err := q.Back()
			if err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
			assert.EqualValues(t, tt.expected.item, item)
		})
	}
}
