package jianzhi_offer

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

func LengthOfLongestSubstring(s string) int {
	// 1.哈希表
	// return Baoli(s)

	// 2.滑动窗口
	return SlideWindow(s)

}

// 滑动窗口
func SlideWindow(s string) int {
	// abcabcbb
	left := 0
	has := make(map[byte]bool)
	max := 0
	for right := 0; right < len(s); right++ {
		for has[s[right]] {
			_, ok := has[s[left]]
			if ok {
				has[s[left]] = false
			}
			left++
		}
		has[s[right]] = true
		max = MaxNum(max, right-left+1)
	}
	return max
}

// 哈希表
func Baoli(s string) int {
	var has map[byte]bool
	max, preMax := 0, 0
	for i := 0; i < len(s); i++ {
		preMax, max = max, 0
		has = make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if has[s[j]] {
				max = j - i
				break
			}
			max++
			has[s[j]] = true
		}
		if preMax > max {
			max = preMax
		}
	}
	return max
}
