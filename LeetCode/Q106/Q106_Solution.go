package Q106


/*
递归 --- 从中序和后序遍历序列构造二叉树
*/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从后序遍历中找到当前子树的根节点
// 在中序遍历中找到该节点的位置，以此位置进行左右子树的分割
// 左右子树 递归

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode { // 构造子树
		if inorderLeft > inorderRight { // 边界条件
			return nil
		}

		val := postorder[len(postorder)-1]       // 后序遍历的最后一个元素即为当前子树的根节点
		postorder = postorder[:len(postorder)-1] // 更新后序遍历
		root := &TreeNode{Val: val}

		inorderRootIndex := idxMap[val]                      // 找到val在中序遍历中的位置，分割为左右子树
		root.Right = build(inorderRootIndex+1, inorderRight) // 递归root的右子树
		root.Left = build(inorderLeft, inorderRootIndex-1)   // 递归root的左子树
		return root
	}
	return build(0, len(inorder)-1)
}
