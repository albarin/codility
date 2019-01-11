package maxprofit

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{}, 0},
		{[]int{5}, 0},
		{[]int{5, 5}, 0},
		{[]int{7, 10, 12, 0, 13}, 13},
		{[]int{23171, 21011, 21123, 21366, 21013, 21367}, 356},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
