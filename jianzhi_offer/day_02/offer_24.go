package day02

// Reference：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

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
