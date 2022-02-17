package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer58ii(b *testing.B) {
	s := "abcdefg"
	k := 2
	rS := ReverseLeftWords(s, k)
	fmt.Println("剑指 Offer 58 - II. 左旋转字符串", rS)
}
