package jianzhi_offer

// Reference：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

func reversePrint(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
