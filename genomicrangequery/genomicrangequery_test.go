package genomicrangequery

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		s string
		p []int
		q []int
		e []int
	}{
		{"C", []int{0}, []int{0}, []int{2}},
		{"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 5}, []int{2, 4, 1}},
		{"AAAAAA", []int{0, 1, 2, 3}, []int{2, 3, 4, 5}, []int{1, 1, 1, 1}},
		{"ACGT", []int{0, 1, 2, 3}, []int{0, 1, 2, 3}, []int{1, 2, 3, 4}},
		{"TCTCTCTCTCTCTTCCTTCTCTC", []int{0, 4}, []int{3, 10}, []int{2, 2}},
		{"AGT", []int{0, 1, 1, 2}, []int{0, 1, 2, 2}, []int{1, 3, 3, 4}},
		{"CC", []int{0, 1, 1}, []int{0, 1, 1}, []int{2, 2, 2}},
		{"AC", []int{0, 0, 1}, []int{0, 1, 1}, []int{1, 1, 2}},
		{"TAT", []int{1}, []int{2}, []int{1}},
	}

	for _, test := range tests {
		e := Solution(test.s, test.p, test.q)
		if !reflect.DeepEqual(e, test.e) {
			t.Errorf("Solution on %s was incorrect, got: %d, expected: %d", test.s, e, test.e)
		}
	}
}
