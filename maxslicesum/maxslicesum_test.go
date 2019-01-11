package maxslicesum

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{5}, 5},
		{[]int{-10}, -10},
		{[]int{-2, -2}, -2},
		{[]int{-6, 6}, 6},
		{[]int{-2, 1, 1}, 2},
		{[]int{3, 2, -6, 4, 0}, 5},
		{[]int{5, -7, 3, 5, -2, 4, -1}, 10},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
