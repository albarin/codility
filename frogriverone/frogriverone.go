package frogriverone

func Solution(X int, A []int) int {
	pos := make([]bool, X)
	numCovered := 0

	for i, n := range A {
		if n <= X && !pos[n-1] {
			pos[n-1] = true
			numCovered++
		}

		if numCovered == X {
			return i
		}
	}

	return -1
}
