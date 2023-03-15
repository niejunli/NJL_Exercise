package Q221

/*
动态规划 --- 最大正方形
能够延展的边长的值是通过left, above,leftAbove来决定的：example
绿色正方形(dp[i-1][j])表示：最多可以为以dp[i][j]为右下角的正方形的右边提供边长为2的边。
黄色正方形(dp[i][j-1])表示：最多可以为以dp[i][j]为右下角的正方形的下边提供边长为3的边。
蓝色正方形(dp[i-1][j-1])表示：最多可以为以dp[i][j]为右下角的正方形的右边和下边提供边长为4的边。
由于题目中说是最大正方形，所以取三者中的最小值(min(dp[i-1][j-1], dp[i-1][j],dp[i][j-1]))是因为根据木桶原理，我们只能取最短的一个边来作为dp[i][j]表示的最大的正方形的边。
加上matrix[i][j]本身的这个正方形的边长1，就得到了以dp[i][j]为右下角最大的正方形的边长。
*/

func maximalSquare(matrix [][]byte) int {
	columns := len(matrix[0]) // 列数

	// 构造dp数组，存储以(i,j)为右下角的最大正方形的边长
	var dp [][]int
	for i := 0; i < len(matrix); i++ {
		dp = append(dp, make([]int, columns))
	}

	// 结果的正方形边长
	var maxSide int

	// 遍历数组，为dp数组赋值
	for i := range matrix {
		for j := range matrix[i] {
			// 为0的话，则不能构成全为1的正方形
			if matrix[i][j] == '0' {
				continue // dp数组元素默认值为0
			}
			// 边界值：第一行 第一列
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				left, above, leftAbove := dp[i][j-1], dp[i-1][j], dp[i-1][j-1]
				dp[i][j] = minFunc(minFunc(left, above), leftAbove) + 1
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}
	return b
}
