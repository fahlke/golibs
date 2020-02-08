package util

// Min finds the minimum value of an arbitrary number of signed integers.
// It returns the minimum signed integer found.
// Time complexity is O(N).
func Min(a int, b ...int) int {
	min := a
	for _, i := range b {
		if i < min {
			min = i
		}
	}

	return min
}
