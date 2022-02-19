package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer62(b *testing.B) {
	n, m := 10, 17
	num := LastRemaining(n, m)
	log.Println("剑指 Offer 62. 圆圈中最后剩下的数字:", num)
}
