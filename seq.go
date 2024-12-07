// Package seq defines various functions useful for sequences and iteration over slices.
package seq

import "iter"

// Window returns a iterator over consecutive sliding window sub-slices of size n over slice s.
// All sub-slices will have length n.
// All sub-slices are clipped to have no capacity beyond the length.
// If s is empty, the sequence is empty: there is no empty slice in the sequence.
// If len(s) < n, the sequence is empty: there is no slice of length n in the sequence.
// Window panics if n is less than 1.
func Window[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice] {
	if n < 1 {
		panic("cannot be less than 1")
	}

	return func(yield func(Slice) bool) {
		for i := n; i <= len(s); i++ {
			start := i - n
			// Set the capacity of each window so that appending to a window does
			// not modify the original slice.
			if !yield(s[start:i:i]) {
				return
			}
		}
	}
}
