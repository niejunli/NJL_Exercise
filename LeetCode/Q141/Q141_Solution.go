package Q141

/*
   快慢指针 --- 判断链表中是否存在环
   用两个指针，一个快指针，一个慢指针，快指针相较于慢指针前进速度快一些
   如果不含有环，快指针最终会遇到null，说明链表不含环；如果含有环，快指针最终会超慢指针一圈，和慢指针相遇，说明链表含有环
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s := head
	q := head
	for q != nil && q.Next != nil {
		s = s.Next
		q = q.Next.Next
		if s == q {
			return true
		}
	}
	return false
}
