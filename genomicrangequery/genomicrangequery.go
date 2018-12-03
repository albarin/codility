package genomicrangequery

import (
	"fmt"
	"math"
)

func Solution(S string, P []int, Q []int) []int {
	M := make([]int, len(P))
	nucs := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	sums := initPrefixSums(S, nucs)

	for i, s := range S {
		for j := range sums {
			sums[j][i+1] = sums[j][i]
		}
		sums[s][i+1]++
	}

	for i := range P {
		p := P[i]
		q := Q[i]

		min := math.MaxInt64
		for j, n := range nucs {
			fmt.Println(string(j), sums[j][q+1]-sums[j][p] > 0, n < min)
			if sums[j][q+1]-sums[j][p] > 0 && n < min {
				M[i] = n
				min = n
			}
		}
	}

	return M
}

func initPrefixSums(S string, nucs map[rune]int) map[rune][]int {
	sums := make(map[rune][]int)

	for nuc := range nucs {
		sums[rune(nuc)] = make([]int, len(S)+1)
	}

	return sums
}
