package oddoccurrenciesinarray

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		o int
	}{
		{[]int{5}, 5},
		{[]int{8, 3, 8}, 3},
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	}

	for _, test := range tests {
		o := Solution(test.a)
		if o != test.o {
			t.Errorf("Odd occurrence on %d was incorrect, got: %d, expected: %d", test.a, o, test.o)
		}
	}
}
