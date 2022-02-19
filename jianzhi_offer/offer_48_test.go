package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer48(b *testing.B) {
	s := "abcabcbb"
	n := LengthOfLongestSubstring(s)
	log.Println("剑指 Offer 48. 最长不含重复字符的子字符串:", n)
}
