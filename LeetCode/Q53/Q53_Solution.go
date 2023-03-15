package Q53

/*
动态规划 --- 最大子数组和
*/

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	maxValue := nums[0]
	maxDP := make([]int, length)
	maxDP[0] = nums[0]
	for i := 1; i < length; i++ {
		maxDP[i] = Max(maxDP[i-1]+nums[i], nums[i])
		maxValue = Max(maxValue, maxDP[i])
	}
	return maxValue
}

func maxSubArray2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < length; i++ {
		currentSum = Max(nums[i], currentSum+nums[i])
		maxSum = Max(maxSum, currentSum)
	}
	return maxSum
}
