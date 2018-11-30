package tapeequilibrium

import (
	"math"
)

func Solution(A []int) int {
	P := len(A) - 1

	var left, right int
	for _, a := range A {
		right += a
	}

	minDiff := math.MaxInt64
	for i := 0; i <= P-1; i++ {
		left += A[i]
		right -= A[i]

		diff := int(math.Abs(float64(left - right)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
