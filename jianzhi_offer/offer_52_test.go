package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer52(b *testing.B) {
	headA := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
	}
	headB := &ListNode{
		Val: 6, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
	}

	fistNode := GetIntersectionNode(headA, headB)
	fmt.Printf("两个链表的第一个公共节点:%v", fistNode)
	fmt.Println()
}
