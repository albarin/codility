package template

func Solution(A []int) int {
	perm := make([]bool, len(A))
	for _, e := range A {
		if e > len(A) {
			return 0
		}

		perm[e-1] = true
	}

	for _, e := range perm {
		if !e {
			return 0
		}
	}

	return 1
}

//func Solution(A []int) int {
//	N := len(A)
//	sum := N * (N + 1) / 2
//
//	var total int
//	for _, e := range A {
//		total += e
//	}
//
//	if sum == total {
//		return 1
//	}
//
//	return 0
//}
