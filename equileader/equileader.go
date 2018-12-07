package equileader

func Solution(A []int) int {
	total := 0
	leader := getLeader(A)
	if leader == -1 {
		return 0
	}

	leaderSum := make([]int, len(A)+1)

	for i := 0; i < len(A); i++ {
		leaderSum[i+1] = leaderSum[i]
		if A[i] == leader {
			leaderSum[i+1]++
		}
	}

	N := len(A)
	for i := 1; i < N; i++ {
		leftSize := i
		rightSize := N - leftSize

		countLeft := leaderSum[i]
		countRight := leaderSum[N] - countLeft

		if countLeft > leftSize/2 && countRight > rightSize/2 {
			total++
		}
	}

	return total
}

func getLeader(A []int) int {
	leaderCount := make(map[int]int)

	currLeader := 0
	for _, e := range A {
		leaderCount[e]++
		if leaderCount[e] > leaderCount[currLeader] {
			currLeader = e
		}
	}

	if leaderCount[currLeader] > len(A)/2 {
		return currLeader
	}

	return -1
}
