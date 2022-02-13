package jianzhi_offer

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

func ValidateStackSequences(pushed []int, popped []int) bool {
	// 构建辅助栈
	stack := make([]int, 0)

	// 弹出哪个值
	i := 0

	// 压入序列遍历
	for j := 0; j < len(pushed); j++ {
		stack = append(stack, pushed[j])
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}

	// 辅助栈为空，返回true
	return len(stack) == 0
}
