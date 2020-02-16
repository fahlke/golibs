package queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_PushEmpty(t *testing.T) {
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

func TestQueue_PushSize(t *testing.T) {
	t.Parallel()

	golden := []struct {
		name  string
		items []string
		size  int
	}{
		{"empty",
			[]string{},
			0x0,
		},
		{"one item",
			[]string{"foo"},
			0x1,
		},
		{"two items",
			[]string{"foo", "bar"},
			0x2,
		},
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

func TestQueue_PushPop(t *testing.T) {
	t.Parallel()

	type expected struct {
		item interface{}
		err  error
	}

	golden := []struct {
		name     string
		items    []string
		expected expected
		size     int
	}{
		{"empty",
			[]string{},
			expected{nil, errors.New("empty queue")},
			0x0,
		},
		{"one item",
			[]string{"foo"},
			expected{"foo", nil},
			0x0,
		},
		{"two items",
			[]string{"foo", "bar"},
			expected{"foo", nil},
			0x1,
		},
		{"three items",
			[]string{"foo", "bar", "baz"},
			expected{"foo", nil},
			0x2,
		},
	}

	for _, tt := range golden {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := Queue{}

			for _, item := range tt.items {
				q.Push(item)
			}

			item, err := q.Pop()
			if err != nil {
				assert.Equal(t, tt.expected.err, err)
			}
			assert.EqualValues(t, tt.expected.item, item)
			assert.EqualValues(t, tt.size, len(q.data))
		})
	}
}

func TestQueue_PushFrontBack(t *testing.T) { //nolint:funlen
	t.Parallel()

	type expected struct {
		item interface{}
		err  error
	}

	golden := []struct {
		name          string
		items         []string
		expectedFront expected
		expectedBack  expected
		size          int
	}{
		{"empty",
			[]string{},
			expected{nil, errors.New("empty queue")},
			expected{nil, errors.New("empty queue")},
			0x0,
		},
		{"one item",
			[]string{"foo"},
			expected{"foo", nil},
			expected{"foo", nil},
			0x1,
		},
		{"two items",
			[]string{"foo", "bar"},
			expected{"foo", nil},
			expected{"bar", nil},
			0x2,
		},
		{"three items",
			[]string{"foo", "bar", "baz"},
			expected{"foo", nil},
			expected{"baz", nil},
			0x3,
		},
	}

	for _, tt := range golden {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := Queue{}

			for _, item := range tt.items {
				q.Push(item)
			}

			actualFront, err := q.Front()
			if err != nil {
				assert.Equal(t, tt.expectedFront.err, err)
			}
			assert.EqualValues(t, tt.expectedFront.item, actualFront)
			assert.EqualValues(t, tt.size, len(q.data))

			actualBack, err := q.Back()
			if err != nil {
				assert.Equal(t, tt.expectedBack.err, err)
			}
			assert.EqualValues(t, tt.expectedBack.item, actualBack)
			assert.EqualValues(t, tt.size, len(q.data))
		})
	}
}
