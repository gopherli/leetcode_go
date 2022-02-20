package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer60(b *testing.B) {
	n := 2
	dp := DicesProbability(n)
	log.Println("剑指 Offer 60. n个骰子的点数:", dp)
}
