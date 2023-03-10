package tasks

func rob2(nums []int) int {
	ln := len(nums)
	dp := make([]int, ln+1)
	dp[0] = 0
	dp[1] = nums[ln-1]

	for i := 2; i < len(dp); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[ln]
}
