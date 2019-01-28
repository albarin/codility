package ladder

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a []int
		b []int
		e []int
	}{
		{[]int{4, 4, 5, 5, 1}, []int{3, 2, 4, 3, 1}, []int{5, 1, 8, 0, 1}},
		{[]int{2, 4, 5, 5, 1}, []int{3, 2, 4, 5, 1}, []int{2, 1, 8, 8, 1}},
		{[]int{1, 1}, []int{1, 10}, []int{1, 1}},
	}

	for _, test := range tests {
		e := Solution(test.a, test.b)
		if !reflect.DeepEqual(e, test.e) {
			t.Errorf("Solution on (%d, %d) was incorrect, got: %d, expected: %d", test.a, test.b, e, test.e)
		}
	}
}
