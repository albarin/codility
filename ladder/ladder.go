package ladder

import (
	"math"
)

func Solution(A []int, B []int) []int {
	L := len(A)
	res := make([]int, L)

	fib := fibonacci(L)
	for i := range A {
		res[i] = fib[A[i]+1] % int(math.Pow(2.0, float64(B[i])))
	}

	return res
}

func fibonacci(L int) []int {
	fib := make([]int, L+2)
	fib[1] = 1

	for i := 2; i < L+2; i++ {
		fib[i] = (fib[i-1] + fib[i-2]) % int(math.Pow(2.0, 30.0))
	}

	return fib
}
