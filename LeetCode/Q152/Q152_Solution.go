package Q152

/*
动态规划 --- 乘积最大子数组
1. 确定dp数组：dp[i]表示nums数组在[0,i]范围内(连续子数组)的最大乘积
2. 确定递推公式：乘积不同于求和，最大乘积不一定就是前面的最大乘积*nums[i]，还有可能是前面的最小乘积*nums[i]，因为最小乘积可能是一个绝对值很大的负数，如果nums[i]也是负数，那么负负得正，有可能就成了最大乘积了。
所以需要分开讨论，如果nums[i]为正数，则希望以i-1结尾的子数组的乘积是一个尽可能大的正数；
反之如果nums[i]为负数，则希望以i-1结尾的子数组的乘积是一个尽可能小(绝对值尽可能大)的负数，这样就得到了本题的递推公式:
maxDP[i] = max(nums[i], max(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]))
minDP[i] = min(nums[i], min(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]))
3 初始化dp数组：由递推公式可知，dp[i]依赖于dp[i-1]，所以dp[0]一定要初始化，显然maxDP[0]和minDP[0]都是nums[0]。
4 确定遍历顺序：由递推公式可知，应该是从前往后正序遍历。
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	maxDP := make([]int, length)
	minDP := make([]int, length)
	maxDP[0], minDP[0] = nums[0], nums[0]
	maxValue := nums[0]
	for i := 1; i < length; i++ {
		maxDP[i] = Max(nums[i], Max(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]))
		minDP[i] = Min(nums[i], Min(maxDP[i-1]*nums[i], minDP[i-1]*nums[i]))
		maxValue = Max(maxValue, maxDP[i])
	}
	return maxValue
}
