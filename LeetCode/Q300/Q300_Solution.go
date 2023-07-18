package Q300

/*
动态规划 --- 最长递增子序列
*/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 // 初始化递归数组，dp[i]表示以nums[i]结尾的最大递增子序列的长度
	}
	max := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
