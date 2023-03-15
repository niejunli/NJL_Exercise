package Q589

/*
N叉树的前序遍历
*/

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res = []int{}
	res = append(res, root.Val)
	for idx := range root.Children {
		res = append(res, preorder(root.Children[idx])...)
	}
	return res
}

// 迭代
func preorderAnother(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res = []int{}
	var stack = []*Node{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]  // 取出栈中最后一个节点
		stack = stack[:len(stack)-1] // 更新栈
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}
