package jianzhi_offer

import (
	"testing"
)

func TestOffer25(t *testing.T) {
	// [1,2,3]
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	// [1,2,3]
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	result := MergeTowList(l1, l2)

	// 打印链表
	var resultArr []int
	for result != nil {
		resultArr = append(resultArr, result.Val)
		result = result.Next
	}

	t.Fatalf("%v", resultArr)
	t.Fatalf("%v", result)
}
