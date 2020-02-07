package maximumtriplet

import (
	"reflect"
	"testing"
)

func TestMaximumTriplet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{"value a is max", []int{70, -10, 30}, 70},
		{"value a is max", []int{3, 2, 1}, 3},
		{"value a is max", []int{-10, -20, -30}, -10},
		{"value b is max", []int{-1, 0, -10}, 0},
		{"value b is max", []int{-70, -10, -30}, -10},
		{"value b is max", []int{7, 64, 43}, 64},
		{"value c is max", []int{1, 2, 3}, 3},
		{"value c is max", []int{1, 2, 3}, 3},
		{"value c is max", []int{-100, 0, 100}, 100},
		{"value c is max", []int{0, 0, 0}, 0},
		{"value c is max", []int{-100, -100, -100}, -100},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := MaximumTriplet(&tt.values[0], &tt.values[1], &tt.values[2])
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got `%d`, expected `%d`", got, tt.expected)
			}
		})
	}
}
