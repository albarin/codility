package template

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
	}{}

	for _, test := range tests {
		e := Solution()
		if e != test.e {
			t.Errorf("on %d was incorrect, got: %d, expected: %d", test.a, e, test.e)
		}
	}
}
