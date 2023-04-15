package Q102

/*
bfs --- 二叉树的层序遍历2
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		tempLength := len(queue)  // 每层节点的个数
		tempRes := make([]int, 0) // 每层节点的结果存放在slice中
		for i := 0; i < tempLength; i++ {
			temp := queue[0]
			tempRes = append(tempRes, temp.Val)
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, tempRes)
	}
	return res
}
