package Q19

/*
	快慢指针 --- 删除倒数第k个节点
	即找到倒数第k+1个节点，将其后继删除
	在对链表进行操作时，一种常用的技巧是添加一个哑节点（dummy node）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{0, head}
	s, q := temp, head
	for ; n > 0 && q != nil; n-- {
		q = q.Next
	}
	for q != nil {
		s = s.Next
		q = q.Next
	}
	s.Next = s.Next.Next
	return temp.Next
}
