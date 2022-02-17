package jianzhi_offer

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

func ReverseLeftWords(s string, n int) string {
	byteS := []byte(s)
	left := byteS[:n]
	right := byteS[n:]
	return string(right) + string(left)
}
