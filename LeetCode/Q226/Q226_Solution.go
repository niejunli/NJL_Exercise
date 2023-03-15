package Q226

/*
递归 --- 翻转二叉树
二叉树题目，第一想到的应该可以用递归来解决
题目的输入和输出，输出的左右子树的位置跟输入正好是相反的，于是可以递归的交换左右子树来完成
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil { // 递归的终止条件
		return root
	}
	root.Left, root.Right = root.Right, root.Left // 将当前节点的左右子树进行交换
	invertTree(root.Left)                         // 递归交换当前节点的 左子树和右子树
	invertTree(root.Right)
	// 函数返回时就表示当前这个节点，以及它的左右子树，都已经交换完了
	return root
}
