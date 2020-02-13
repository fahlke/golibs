package util

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		n    []int
		want int
	}{
		{"one zero", 0, nil, 0},
		{"one value", 30, nil, 30},
		{"two zeros", 0, []int{0}, 0},
		{"two values", 30, []int{40}, 40},
		{"three values (all positive)", 30, []int{40, 50}, 50},
		{"three values (all negative)", -70, []int{-10, -30}, -10},
		{"three values (all zeros)", 0, []int{0, 0}, 0},
		{"three values (positive and negative)", -100, []int{0, 100}, 100},
		{"edge case (max int64)", math.MaxInt64, []int{10, 20}, math.MaxInt64},
		{"edge case (min int64)", math.MinInt64, nil, math.MinInt64},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Max(tt.a, tt.n...)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
