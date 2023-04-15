package offer_55

/*
递归 --- 二叉树的深度
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return 1 + Max(l, r)
}
