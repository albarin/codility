package numberofdiscintersections

func Solution(A []int) int {
	total := 0
	left := make([]int, len(A))
	right := make([]int, len(A))

	for k := range A {
		left[k] = k - A[k]
		right[k] = k + A[k]
	}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if right[i] >= left[j] {
				total++
			}
			if total > 10000000 {
				return -1
			}
		}
	}

	return total
}
