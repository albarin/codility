package equileader

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		e int
	}{
		{[]int{4, 3, 4, 4, 4, 2}, 2},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
