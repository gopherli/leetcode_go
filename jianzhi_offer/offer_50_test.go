package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer50(b *testing.B) {
	s := "agebaccdeff"

	log.Printf("字符串%v第一个只出现一次的字符为:%v", s, string(FirstUniqChar(s)))
}
