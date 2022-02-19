package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer10(b *testing.B) {
	n := 1
	nums := NumWays(n)
	log.Println("剑指 Offer 10- II. 青蛙跳台阶问题:", nums)
}
