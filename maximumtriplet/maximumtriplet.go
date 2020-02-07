// Package maximumtriplet provides a function to find the highest integer in a list of three.
package maximumtriplet

// MaximumTriplet finds the maximum value of three signed integers.
// It returns the maximum signed integer found.
// Time complexity is O(1).
func MaximumTriplet(a, b, c *int) int {
	switch {
	case *a > *b && *a > *c:
		return *a
	case *b > *c:
		return *b
	default:
		return *c
	}
}
