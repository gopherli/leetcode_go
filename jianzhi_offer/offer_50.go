package jianzhi_offer

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

func FirstUniqChar(s string) byte {
	// 特殊处理
	if len(s) == 0 {
		return ' '
	}

	// string转[]byte
	// b := []byte(s)不需要转，range直接可以遍历string为byte

	// 记录元素出现次数
	countData := make([]int, 26)

	// 遍历b,得出每个元素
	for _, v := range s {
		// 记录元素出现的次数
		countData[v-'a']++
	}

	// 再次遍历b,得出第一个只出现一次的元素
	for k, v := range s {
		if countData[v-'a'] == 1 {
			return s[k]
		}
	}

	return ' '
}
