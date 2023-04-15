package Q718

/*
动态规划 --- 最长重复子数组(连续)
https://leetcode.cn/problems/maximum-length-of-repeated-subarray/solution/zhe-yao-jie-shi-ken-ding-jiu-dong-liao-by-hyj8/
*/

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, m+1)
	// 初始化dp数组
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1) 
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // 以nums1[i-1]，nums2[j-1]为结尾的公共数组长度
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

// dp[i][j] 只依赖上一行上一列的对角线的值，所以可以从右上角开始计算
// 一维数组 dp ， dp[j] 是以 A[i-1], B[j-1] 为末尾项的最长公共子数组的长度
func otherFindLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([]int, n+1) // 初始化dp数组
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- { // 从右边开始遍历
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1 // 取决于左上角的值
			} else {
				dp[j] = 0 // 需要覆盖，因为每一行的结果都用这个slice
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}
