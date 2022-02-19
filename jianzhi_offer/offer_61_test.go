package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer61(b *testing.B) {
	nums := []int{2, 2, 3, 0, 0}
	bl := IsStraight(nums)
	log.Printf("扑克牌中的顺子%v是否正确：%v", nums, bl)
}
