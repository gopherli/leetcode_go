package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer57I(b *testing.B) {
	target := 9
	res := FindContinuousSequence(target)
	log.Println("和为s的连续正整数序列：", res)
}
