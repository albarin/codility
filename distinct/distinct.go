package distinct

import (
	"sort"
)

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	total := 1
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			total++
		}
	}

	return total
}
