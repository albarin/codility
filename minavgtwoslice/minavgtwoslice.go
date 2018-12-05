package minavgtwoslice

import (
	"math"
)

func Solution(A []int) int {
	minPos := 0
	min := math.MaxFloat64

	for i := 2; i < len(A); i++ {
		avg01 := float64(A[i-2]+A[i-1]) / 2
		avg02 := float64(A[i-2]+A[i-1]+A[i]) / 3
		avg12 := float64(A[i-1]+A[i]) / 2

		currMin0 := math.Min(avg01, avg02)
		currMin1 := math.Min(currMin0, avg12)

		if currMin0 <= currMin1 && currMin0 < min {
			min = currMin0
			minPos = i - 2
		} else if currMin1 < min {
			min = currMin1
			minPos = i - 1
		}
	}

	return int(minPos)
}
