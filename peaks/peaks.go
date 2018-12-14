package peaks

import (
	"sort"
)

func Solution(A []int) int {
	N := len(A)

	if N < 3 {
		return 0
	}

	total, peaks := getPeaks(A)
	divs := getDivisors(N)

	if total == 0 {
		return 0
	}

	for _, f := range divs {
		hasPeak := true
		for j := 0; j < N; j = j + f {
			if !peakExists(peaks, j, j+f) {
				hasPeak = false
				break
			}
		}

		if hasPeak {
			return N / f
		}
	}

	return 1
}

func peakExists(peaks []int, s int, e int) bool {
	var start, end int
	if s-1 < 0 {
		start = 0
	} else {
		start = peaks[s-1]
	}

	end = peaks[e-1]

	return end-start > 0
}

func getPeaks(A []int) (int, []int) {
	peaks := make([]int, len(A))
	totalPeaks := 0

	for i := 1; i < len(A); i++ {
		peaks[i] = peaks[i-1]
		if i+1 < len(A) && A[i] > A[i-1] && A[i] > A[i+1] {
			totalPeaks++
			peaks[i]++
		}
	}

	return totalPeaks, peaks
}

func getDivisors(N int) []int {
	divs := make([]int, 0)

	i := 2
	for i*i < N {
		if N%i == 0 {
			divs = append(divs, i, N/i)
		}
		i++
	}

	if i*i == N {
		divs = append(divs, i)
	}

	sort.Ints(divs)

	return divs
}
