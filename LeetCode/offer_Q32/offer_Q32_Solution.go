package offer_Q32

/*
bfs --- 二叉树的层次遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0) // bfs利用队列实现 利用先进先出的性质
	queue = append(queue, root)

	for len(queue) > 0 { // 跳出条件：队列为空
		temp := queue[0]            // 取节点
		res = append(res, temp.Val) // 取值
		queue = queue[1:]           // 将节点从队列弹出

		if temp.Left != nil { // 节点的左孩子
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil { // 节点的右孩子
			queue = append(queue, temp.Right)
		}
	}
	return res
}
