package jianzhi_offer

// Reference：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
