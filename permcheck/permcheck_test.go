package template

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		p int
	}{
		{[]int{4}, 0},
		{[]int{1}, 1},
		{[]int{4, 1, 3}, 0},
		{[]int{4, 1, 3, 2}, 1},
		{[]int{4, 1, 3, 2, 4}, 0},
		{[]int{3, 1, 3, 3, 5}, 0},
	}

	for _, test := range tests {
		p := Solution(test.a)
		if p != test.p {
			t.Errorf("Missing element on %d was incorrect, got: %d, expected: %d", test.a, p, test.p)
		}
	}
}
