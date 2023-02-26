package tasks

func findJudge(n int, trust [][]int) int {

	if n == 1 && len(trust) > 0 {
		return -1
	} else if n == 1 {
		return 1
	}
	checkMap := make(map[int]int)
	for _, sl := range trust {
		checkMap[sl[0]]--
		checkMap[sl[1]]++
	}
	for key, val := range checkMap {
		if val == n-1 {
			return key
		}
	}
	return -1
}
