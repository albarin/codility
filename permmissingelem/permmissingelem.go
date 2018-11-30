package permmissingelem

func Solution(A []int) int {
	n := len(A) + 1
	sum := n * (n + 1) / 2

	for _, num := range A {
		sum -= num
	}

	return sum
}
