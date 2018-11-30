package binarygap

import (
	"strconv"
)

func Solution(N int) int {
	binN := strconv.FormatInt(int64(N), 2)

	var longestGap, gap int
	for _, bit := range binN {
		if bit == '1' {
			if gap > longestGap {
				longestGap = gap
			}

			gap = 0
		} else {
			gap++
		}
	}

	return longestGap
}
