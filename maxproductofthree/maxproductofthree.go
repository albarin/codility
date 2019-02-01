package maxproductofthree

import (
	"math"
	"sort"
)

func Solution(A []int) int {
	N := len(A)
	sort.Ints(A)

	p1 := A[0] * A[1] * A[N-1]
	p2 := A[N-1] * A[N-2] * A[N-3]

	return int(math.Max(float64(p1), float64(p2)))
}
