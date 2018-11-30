package binarygap

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		n int
		gap int
	} {
		{9, 2},
		{529, 4},
		{20, 1},
		{15, 0},
		{32, 0},
		{1041, 5},
	}

	for _, test := range tests {
		longestGap := Solution(test.n)
		if longestGap != test.gap {
			t.Errorf("Longest binary gap of %d was incorrect, got: %d, expected: %d", test.n, longestGap, test.gap)
		}
	}
}