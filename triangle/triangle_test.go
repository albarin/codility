package triangle

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{10, 2, 5, 1, 8, 20}, 1},
		{[]int{10, 50, 5, 1}, 0},
		{[]int{5, 3, 3}, 1},
		{[]int{-100, 2, 4, 5}, 1},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
