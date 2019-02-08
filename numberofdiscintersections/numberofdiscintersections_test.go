package numberofdiscintersections

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{0}, 0},
		{[]int{4}, 0},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 5, 2, 1, 4, 0}, 11},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
