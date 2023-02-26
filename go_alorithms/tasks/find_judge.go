package tasks

func findJudge(n int, trust [][]int) int {
	mostTrusted := -1
	cntMap := make([]int, n)

	for _, t := range trust {
		cntMap[t[0]]++
		mostTrusted = t[1]
	}
}
