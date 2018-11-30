package permmissingelem

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{}, 1},
		{[]int{1}, 2},
		{[]int{2, 3, 1, 5}, 4},
		{[]int{8, 4, 2, 1, 5, 6, 7, 3}, 9},
		{[]int{8, 4, 2, 9, 5, 6, 7, 3}, 1},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Missing element on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
