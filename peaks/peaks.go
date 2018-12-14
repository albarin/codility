package peaks

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	N := len(A)
	peaks := getPeaks(A)
	divs := getDivisors(N)

	fmt.Println("--------------------")
	fmt.Println("peaks:", peaks)
	fmt.Println("factors:", divs)

	if len(peaks) == 0 {
		return 0
	}

	for _, f := range divs {
		fmt.Println("factor:", f)
		hasPeak := true
		for j := 0; j < N; j = j + f {
			fmt.Println("block:", j, j+f-1)
			if !peakExists(peaks, j, j+f, N) {
				fmt.Println("    no peak in block")
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

func peakExists(peaks []int, start int, end int, N int) bool {
	fmt.Println("checking peaks between:", start, end)
	for i, p := range peaks {
		if i > end {
			break
		}
		fmt.Printf("p=%d, start=%d, end=%d, N=%d\n", p, start, end, N)
		if (p >= start) && (p < end) {
			fmt.Println("    peak was found")
			return true
		}
	}

	return false
}

func getPeaks(A []int) []int {
	peaks := make([]int, 0)

	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	return peaks
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
