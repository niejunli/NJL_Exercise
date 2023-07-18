package offer_Q25

/*
虚拟头节点+递归 --- 合并有序链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var n *ListNode
	if l1.Val < l2.Val {
		n = l1
		n.Next = mergeTwoLists(l1.Next, l2)
	} else {
		n = l2
		n.Next = mergeTwoLists(l1, l2.Next)
	}
	return n
}
