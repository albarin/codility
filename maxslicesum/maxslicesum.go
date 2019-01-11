package maxslicesum

func Solution(A []int) int {
	m := A[0]
	sum := A[0]

	for i := 1; i < len(A); i++ {
		if sum+A[i] > A[i] {
			sum = sum + A[i]
		} else {
			sum = A[i]
		}

		if sum > m {
			m = sum
		}
	}

	return m
}
