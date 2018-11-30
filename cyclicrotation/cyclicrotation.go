package cyclicrotation

func Solution(A []int, K int) []int {
	lenA := len(A)
	if K == 0 || K == lenA {
		return A
	}

	var rotIndex int
	R := make([]int, lenA)
	for i, n := range A {
		rotIndex = i + K
		if i+K >= lenA {
			rotIndex = (i + K) % lenA
		}

		R[rotIndex] = n
	}

	return R
}
