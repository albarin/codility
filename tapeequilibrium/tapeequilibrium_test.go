package tapeequilibrium

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		d int
	}{
		{[]int{3, 3}, 0},
		{[]int{1, -1}, 2},
		{[]int{-1, -100}, 99},
		{[]int{3, 1, 2, 4, 3}, 1},
	}

	for _, test := range tests {
		d := Solution(test.a)
		if d != test.d {
			t.Errorf("Wrong minimal difference for %d, got: %d, expected: %d", test.a, d, test.d)
		}
	}
}
