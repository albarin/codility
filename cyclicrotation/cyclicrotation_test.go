package cyclicrotation

import (
	"reflect"
	"testing"
)

func TestSolution_RotationTimeEqualsArraySize(t *testing.T) {
	input := []int{1, 2, 3, 4}
	solution := Solution(input, 4)
	if !reflect.DeepEqual(solution, input) {
		t.Errorf("Rotation of %d was incorrect, got: %d, expected: %d", input, solution, input)
	}
}

func TestSolution(t *testing.T) {
	tests := []struct {
		a        []int
		k        int
		rotation []int
	}{
		{[]int{2, 8, 9, 7, 6}, 1, []int{6, 2, 8, 9, 7}},
		{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{[]int{6, 4, 3, 1, 0}, 4, []int{4, 3, 1, 0, 6}},
		{[]int{0, 0, 0, 0}, 1, []int{0, 0, 0, 0}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		solution := Solution(test.a, test.k)
		if !reflect.DeepEqual(solution, test.rotation) {
			t.Errorf("Rotation of %d was incorrect, got: %d, expected: %d", test.a, solution, test.a)
		}
	}
}
