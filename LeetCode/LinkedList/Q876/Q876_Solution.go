package Q876

/*
	快慢指针 --- 链表中间节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s, q := head, head
	for q != nil && q.Next != nil {
		s = s.Next
		q = q.Next.Next
	}
	return s
}
