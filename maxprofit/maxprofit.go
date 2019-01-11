package maxprofit

func Solution(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	max := 0
	start := A[0]

	for i := 1; i < len(A); i++ {
		diff := A[i] - start
		if A[i] > start {
			if diff > max {
				max = diff
			}
		} else {
			start = A[i]
		}
	}

	return max
}
