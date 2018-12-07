package equileader

func Solution(A []int) int {
	equileaders := 0
	leader, count := getLeader(A)
	if leader == -1 {
		return 0
	}

	N := len(A)
	leftCount := 0
	for i, e := range A {
		if e == leader {
			leftCount++
		}
		leftSize := i + 1
		rightSize := N - leftSize

		rightCount := count - leftCount

		if leftCount > leftSize/2 && rightCount > rightSize/2 {
			equileaders++
		}
	}

	return equileaders
}

func getLeader(A []int) (int, int) {
	leaderCount := make(map[int]int)

	currLeader := 0
	for _, e := range A {
		leaderCount[e]++
		if leaderCount[e] > leaderCount[currLeader] {
			currLeader = e
		}
	}

	if leaderCount[currLeader] > len(A)/2 {
		return currLeader, leaderCount[currLeader]
	}

	return -1, 0
}
