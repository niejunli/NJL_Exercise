package Q200

/*
dfs深度优先搜索 --- 岛屿数量
*/

func numIslands(grid [][]byte) int {
	var ans int
	m, n := len(grid), len(grid[0]) // m行 n列
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return // bad case，递归终止条件：越界(无效的网格)
		}
		if grid[r][c] != '1' {
			return // 海洋/已访问过的陆地
		}
		grid[r][c] = '2' // 标记已访问的陆地，防止无限循环无法退出
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' { // 遇到一个陆地，结果+1，并将其周围四个相连的陆地变成非陆地
				ans++
				dfs(r, c)
			}
		}
	}
	return ans
}
