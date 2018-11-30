package maxcounters

func Solution(N int, A []int) []int {
	maxCounter := 0
	lastMax := 0
	counters := make([]int, N)

	for _, op := range A {
		if op <= N {
			if counters[op-1] < lastMax {
				counters[op-1] = lastMax + 1
			} else {
				counters[op-1]++
			}

			if counters[op-1] > maxCounter {
				maxCounter = counters[op-1]
			}
		} else {
			lastMax = maxCounter
		}
	}

	for i := range counters {
		if counters[i] < lastMax {
			counters[i] = lastMax
		}
	}

	return counters
}
