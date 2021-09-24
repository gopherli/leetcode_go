package day02

import (
	"fmt"
	"testing"
)

func TestOffer06(t *testing.T) {
	head := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	heads := ListNode{}
	fmt.Print(heads)
	revArr := reversePrint(&head)
	fmt.Println(revArr)
}
