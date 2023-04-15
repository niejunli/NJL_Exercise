package Q46

/*
回溯 --- 全排列
*/

func permute(nums []int) [][]int {
	var res, path = make([][]int, 0), make([]int, 0) // 注意初始化这里不要写成make([]int, len(nums))
	var used = make([]bool, len(nums))               // 注意不要写成make([]bool, 0)

	var dfs func()
	dfs = func() {
		// 其中一个分支满足条件
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := range nums {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			//  一个递归调用结束，结束当前分支。回到选择前的状态，去别的分支继续搜
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
