package day02

// Reference：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := []int{}

	if head == nil {
		return arr
	}

	for head.Next != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	arr = append(arr, head.Val)

	long := len(arr)
	revArr := make([]int, long)
	for i := 0; i < len(arr); i++ {
		long--
		revArr[long] = arr[i]
	}
	return revArr
}
