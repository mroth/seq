package seq

import (
	"slices"
	"testing"
)

func TestWindow(t *testing.T) {
	cases := []struct {
		name    string
		s       []int
		n       int
		windows [][]int
	}{
		{
			name:    "nil",
			s:       nil,
			n:       1,
			windows: nil,
		},
		{
			name:    "empty",
			s:       []int{},
			n:       1,
			windows: nil,
		},
		{
			name:    "short",
			s:       []int{1, 2},
			n:       3,
			windows: nil,
		},
		{
			name:    "one",
			s:       []int{1, 2},
			n:       2,
			windows: [][]int{{1, 2}},
		},
		{
			name:    "even",
			s:       []int{1, 2, 3, 4},
			n:       2,
			windows: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name:    "odd",
			s:       []int{1, 2, 3, 4, 5},
			n:       2,
			windows: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var windows [][]int
			for c := range Window(tc.s, tc.n) {
				windows = append(windows, c)
			}

			if !slices.EqualFunc(windows, tc.windows, slices.Equal) {
				t.Errorf("Window(%v, %d) = %v, want %v", tc.s, tc.n, windows, tc.windows)
			}

			if len(windows) == 0 {
				return
			}

			// Verify that appending to the end of the first window does not
			// clobber the beginning of the next window.
			s := slices.Clone(tc.s)
			windows[0] = append(windows[0], -1)
			if !slices.Equal(s, tc.s) {
				t.Errorf("slice was clobbered: %v, want %v", s, tc.s)
			}
		})
	}
}
