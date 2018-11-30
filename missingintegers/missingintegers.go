package missingintegers

func Solution(A []int) int {
	numbers := make(map[int]bool)
	for _, n := range A {
		if n > 0 {
			numbers[n] = true
		}
	}

	i := 1
	for {
		if !numbers[i] {
			return i
		}

		i++
	}

	return i
}
