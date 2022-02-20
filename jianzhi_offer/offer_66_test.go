package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer66(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	res := ConstructArr(a)
	log.Println("剑指 Offer 66. 构建乘积数组:", res)
}
