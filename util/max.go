package util

// Max finds the maximum value of an arbitrary number of signed integers.
// It returns the maximum signed integer found.
// Time complexity is O(N).
func Max(a int, b ...int) int {
	max := a
	for _, i := range b {
		if i > max {
			max = i
		}
	}

	return max
}
