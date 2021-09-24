package day02

import (
	"fmt"
	"testing"
)

func TestOffer24(t *testing.T) {
	head := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	revList := reverseList(&head)
	fmt.Print(revList.Next.Next)
}
