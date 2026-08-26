package townjudge

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	inD := make([]int, n+1)
	oD := make([]int, n+1)

	for _, t := range trust {
		inD[t[1]]++
		oD[t[0]]++
	}

	i := 0
	for i < len(inD) {
		if inD[i] == (n-1) && oD[i] == 0 {
			return i
		}
		i++
	}

	return -1
}
