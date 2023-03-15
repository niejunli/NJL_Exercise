package q695

/*
dfs --- 岛屿的最大面积
*/

var (
	// 上 下 左 右 四个方向
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, r, c int) (area int) {
	if grid[r][c] == 1 {
		grid[r][c] = 2 // 标记一下，防止重复遍历
		area++
		for k := 0; k < 4; k++ {
			mx, my := r+dx[k], c+dy[k]
			if mx >= 0 && mx < len(grid) && my >= 0 && my < len(grid[0]) { // 非bad case
				area += dfs(grid, mx, my)
			}
		}
	}
	return
}

func maxAreaOfIsland(grid [][]int) int {
	var ans = 0
	// 逐个元素遍历
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 { // 碰到1说明是陆地
				ans = Max(ans, dfs(grid, r, c)) // dfs(grid, r, c)以(r,c)为起点开始搜索
			}
		}
	}
	return ans
}
