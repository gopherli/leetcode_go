package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer49(b *testing.B) {
	n := 10
	num := NthUglyNumber(n)
	log.Println("εζ Offer 49. δΈζ°:", num)
}
