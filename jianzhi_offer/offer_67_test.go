package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer67(b *testing.B) {
	str := "   42 "
	num := StrToInt(str)
	log.Println("剑指 Offer 67. 把字符串转换成整数:", num)
}
