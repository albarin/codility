package missingintegers

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{-1, -3}, 1},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 3, 6, 4, 1, 2}, 5},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
