package Q144

/*
二叉树的前序遍历
递归：函数自己调用自己
迭代：A重复调用B
*/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res = []int{}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}


// 迭代
func preorderTraversalAnother(root *TreeNode) []int {
    var res []int
    var stack []*TreeNode
    for root != nil || len(stack) > 0 {
        for root != nil { 
            res = append(res, root.Val) // 根节点放入结果
            stack = append(stack, root.Right) // 右节点暂存（可能为nil）
            root = root.Left // 指向左节点
        }
        top := len(stack)-1 // 挨个弹出右节点
        root = stack[top]
        stack = stack[:top]
    }
    return res
}