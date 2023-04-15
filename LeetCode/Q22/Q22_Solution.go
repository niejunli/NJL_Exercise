package Q22

/*
 dfs --- 括号生成
 https://leetcode.cn/problems/generate-parentheses/solution/shen-du-you-xian-bian-li-zui-jian-jie-yi-ypti/
括号生效的条件：1. 左右括号数量相等 2. 任何前缀中，左括号的数量>=右括号的数量
dfs搜索函数设计： 1. 当前位的选择是左括号还是右括号 2. 如果左括号数量不大于n,可以选择放左括号；如果右括号数量不大于n且右括号数量<左括号数量,可以选择放右括号
*/

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	dfs(n, 0, 0, "")
	return res
}

func dfs(n, lc, rc int, path string) { // n为括号对数，lc是左括号数量，rc是右括号数量，path是当前维护的合法括号序列
	if lc == n && rc == n {
		res = append(res, path)
	} else {
		if lc < n {
			dfs(n, lc+1, rc, path+"(") // 放置左括号
		}
		if rc < lc && rc < n {
			dfs(n, lc, rc+1, path+")") // 放置右括号
		}
	}
}
