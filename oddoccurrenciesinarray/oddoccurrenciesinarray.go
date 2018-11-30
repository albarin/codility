package oddoccurrenciesinarray

func Solution(A []int) int {
	occ := make(map[int]int)

	for _, n := range A {
		occ[n]++
	}

	for i, n := range occ {
		if n%2 != 0 {
			return i
		}
	}

	return 0
}
