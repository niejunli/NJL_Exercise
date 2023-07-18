package Q206

/*
反转链表
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode // 当前节点反转后指向的节点
	cur := head       // 当前节点
	for cur != nil {
		next := cur.Next // 原链表当前节点指向的节点
		cur.Next = pre   // 反转
		pre = cur
		cur = next
	}
	return pre
}
