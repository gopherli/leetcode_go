package jianzhi_offer

// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

func ReverseWords(s string) string {
	// 去掉空格
	moveSlideEmpty := func(s string) []byte {
		left, right := 0, len(s)-1
		// 去掉头空格
		for left <= right && s[left] == ' ' {
			left++
		}

		// 去掉尾空格
		for left <= right && s[right] == ' ' {
			right--
		}

		// 去掉中间空格
		sb := make([]byte, 0)
		for left <= right {
			if s[left] != ' ' {
				sb = append(sb, s[left])
			} else if s[len(sb)-1] != ' ' {
				sb = append(sb, s[left])
			}
			left++
		}
		return sb
	}

	// 翻转字符串
	reverseString := func(sb []byte, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			sb[left], sb[right] = sb[right], sb[left]
		}
	}

	// 翻转字母
	reverseWords := func(sb []byte) {
		start, end, n := 0, 0, len(sb)-1
		for start <= n {
			for end <= n && sb[end] != ' ' {
				end++
			}
			reverseString(sb, start, end-1)
			start, end = end+1, end+1
		}
	}

	sb := moveSlideEmpty(s)
	reverseString(sb, 0, len(sb)-1)
	reverseWords(sb)

	return string(sb)
}
