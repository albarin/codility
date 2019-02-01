package triangle

import (
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)

	for i := 0; i < len(A)-2; i++ {
		c1 := A[i]+A[i+1] > A[i+2]
		c2 := A[i+1]+A[i+2] > A[i]
		c3 := A[i]+A[i+2] > A[i+1]

		if c1 && c2 && c3 {
			return 1
		}
	}

	return 0
}
