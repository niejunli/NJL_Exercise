package Q142

/*
	快慢指针 --- 判断链表中环开始的节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	s, q := head, head
	for q != nil && q.Next != nil {
		s = s.Next
		q = q.Next.Next
		if s == q { // 说明链表中存在环
			s = head     // 将其中一个指针重新指向头指针
			for s != q { // 再次相遇时即为环开始的节点
				s = s.Next
				q = q.Next
			}
			return s
		}
	}
	return nil
}
