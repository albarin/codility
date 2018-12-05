package minavgtwoslice

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		//{[]int{0, 1}, 0},
		//{[]int{8, 1, 1}, 1},
		//{[]int{4, 2, 2, 5, 1, 5, 8}, 1},
		{[]int{-3, -5, -8, -4, -10}, 2},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution for %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
