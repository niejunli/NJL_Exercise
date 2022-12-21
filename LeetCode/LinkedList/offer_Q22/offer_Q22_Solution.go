package offer_Q22

/*
	快慢指针 --- 链表中倒数第k个节点
	构建两个指针的相对距离为k，那么当快指针指向链表末尾时，
	慢指针和快指针的距离还是k，即指向倒数第k个节点
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	s, q := head, head
	for ; k > 0 && q != nil; k-- {
		q = q.Next
	}
	for q != nil {
		s = s.Next
		q = q.Next
	}
	return s
}
