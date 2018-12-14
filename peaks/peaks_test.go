package peaks

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{0, 1, 0, 0, 1, 0, 0, 1, 0}, 3},
		{[]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, 3},
		{[]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2, 7, 3}, 2},
		{[]int{6, 2}, 0},
		{[]int{1, 3, 2, 1}, 1},

		{[]int{1, 3, 2, 4, 3, 4, 1, 2, 3, 4, 6, 2, 7, 3, 1, 1}, 4},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
