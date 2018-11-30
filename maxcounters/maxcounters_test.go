package maxcounters

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		n int
		a []int
		c []int
	}{
		{1, []int{1}, []int{1}},
		{2, []int{1}, []int{1, 0}},
		{5, []int{1, 6}, []int{1, 1, 1, 1, 1}},
		{2, []int{1, 1, 3}, []int{2, 2}},
		{5, []int{3}, []int{0, 0, 1, 0, 0}},
		{5, []int{3, 4}, []int{0, 0, 1, 1, 0}},
		{5, []int{3, 4, 4}, []int{0, 0, 1, 2, 0}},
		{5, []int{3, 4, 4, 6}, []int{2, 2, 2, 2, 2}},
		{5, []int{3, 4, 4, 6, 1}, []int{3, 2, 2, 2, 2}},
		{5, []int{3, 4, 4, 6, 1, 4}, []int{3, 2, 2, 3, 2}},
		{5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2}},
		{5, []int{9, 9, 9, 9, 9, 9, 9, 9, 9}, []int{0, 0, 0, 0, 0}},
		{5, []int{9, 9, 9, 9, 9, 9, 1, 9, 9}, []int{1, 1, 1, 1, 1}},
		{5, []int{3, 4, 4, 6, 1, 4, 4, 6, 6}, []int{4, 4, 4, 4, 4}},
	}

	for _, test := range tests {
		c := Solution(test.n, test.a)
		if !reflect.DeepEqual(c, test.c) {
			t.Errorf("Counters on (%d, %d) were incorrect, got: %d, expected: %d", test.n, test.a, c, test.c)
		}
	}
}
