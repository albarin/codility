package frogjmp

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		x int
		y int
		d int
		j int
	}{
		{19, 85, 30, 3},
		{1, 1, 10, 0},
		{20, 300, 5, 56},
	}

	for _, test := range tests {
		j := Solution(test.x, test.y, test.d)
		if j != test.j {
			t.Errorf("Wrong number on jumps for (%d, %d, %d), got: %d, expected: %d", test.x, test.y, test.d, j, test.j)
		}
	}
}
