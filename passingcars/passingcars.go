package template

func Solution(A []int) int {
	east := 0
	passing := 0

	for _, n := range A {
		if n == 0 {
			east++
		} else {
			passing += east
		}

		if passing > 1e9 {
			return -1
		}
	}

	return passing
}
