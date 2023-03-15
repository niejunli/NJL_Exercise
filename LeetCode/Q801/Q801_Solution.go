package Q801

/*
动态规划 --- 使序列递增的最小交换次数
严格递增：只需要考虑相邻元素的大小关系，即可以从左至右递推计算出答案
状态转移：当前元素的交换与否会影响到后续元素的交换与否，因此需要在下标i的基础上，额外加上一个维度表示是否发生交换
dp[i][0]表示不交换nums1[i]和nums2[i]
dp[i][1]表示交换nums1[i]和nums2[i]
根据相邻的元素分类讨论：
a1 = nums1[i-1]  a2 = nums1[i]
b1 = nums2[i-1]  b2 = nums2[i]
if a1<a2 && b1<b2 then 这两对数可以都交换，也可以都不交换
if a1<b2 && b1<a2 then 可以交换其中一对
如果同时满足以上两种情况，则在转移时取较小值
*/

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSwap1(nums1 []int, nums2 []int) int {
	length := len(nums1)
	dp := make([][2]int, length)
	dp[0][1] = 1 // idx==0位置的元素进行交换，交换次数为1(初始状态)
	for i := 1; i < length; i++ {
		dp[i] = [2]int{length, length} // 答案不会超过 n，故初始化成 n 方便后面取 Min
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			dp[i][0] = dp[i-1][0]     // 可以都不交换，即i-1位置未交换，同样的i位置也未交换
			dp[i][1] = dp[i-1][1] + 1 // 可以都交换
		}
		if nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i] {
			dp[i][0] = Min(dp[i-1][1], dp[i][0]) // 可以交换其中一对；并且同时满足，取较小值
			dp[i][1] = Min(dp[i-1][0]+1, dp[i][1])
		}
	}
	return Min(dp[length-1][0], dp[length-1][1])
}

// 状态转移只发生在i-1和i之间，可以用两个变量来表示状态转移过程
func minSwap2(nums1, nums2 []int) int {
	n, notSwap, swap := len(nums1), 0, 1
	for i := 1; i < n; i++ {
		ns, s := n, n // 答案不会超过 n，故初始化成 n 方便后面取 min
		if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			ns, s = notSwap, swap+1
		}
		if nums2[i-1] < nums1[i] && nums1[i-1] < nums2[i] {
			ns, s = Min(ns, swap), Min(s, notSwap+1)
		}
		notSwap, swap = ns, s
	}
	return Min(notSwap, swap)
}
