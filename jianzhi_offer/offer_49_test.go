package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer49(b *testing.B) {
	n := 10
	num := NthUglyNumber(n)
	log.Println("剑指 Offer 49. 丑数:", num)
}
