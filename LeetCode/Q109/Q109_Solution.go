package Q109

/*
递归 --- 有序链表转换二叉搜索树
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var sortedArray []int
	for head != nil {
		sortedArray = append(sortedArray, head.Val)
		head = head.Next
	}
	var dfs func([]int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		mid := len(nums) / 2 // 为了转化为高度平衡的二叉搜索树(左右子树的高度差<1)  链表本身是升序的，选择中位数作为root，分别递归求左右子树
		root := &TreeNode{nums[mid], nil, nil}
		root.Left = dfs(nums[:mid])
		root.Right = dfs(nums[mid+1:])
		return root
	}
	return dfs(sortedArray)
}
