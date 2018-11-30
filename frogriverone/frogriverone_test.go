package frogriverone

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		x int
		a []int
		t int
	}{
		{1, []int{1}, 0},
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
		{3, []int{1, 3, 1, 4, 2, 3, 5, 4}, 4},
		{1, []int{1, 3, 1, 4, 2, 3, 5, 4}, 0},
		{3, []int{3, 1, 4, 2, 3, 5, 4}, 3},
	}

	for _, test := range tests {
		e := Solution(test.x, test.a)
		if e != test.t {
			t.Errorf("Earliest time on (%d, %d) was incorrect, got: %d, expected: %d", test.x, test.a, e, test.t)
		}
	}
}
