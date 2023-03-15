package offer_Q47

/*
动态规划 --- 礼物的最大价值
对于每一个单元格来说，只能通过其左边的单元格或上边的单元格到达 // 重叠子问题
dp数组保存走到当前grid单元格能够得到的最大价值
dp[i][j] = max(dp[i-1][j],dp[i][j-1])+grid[i][j] -> 状态转移方程
第一行 第一列的处理 -> 边界情况
*/

func maxValue(grid [][]int) int {
	// 处理边界
	// 第一行
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	// 第一列
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	// 其他位置
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = maxFunc(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
