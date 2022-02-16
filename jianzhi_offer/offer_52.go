package jianzhi_offer

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func Hash52(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	has := make(map[*ListNode]bool, 0)
	for headA != nil {
		has[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if has[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
