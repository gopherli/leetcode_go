package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer39(b *testing.B) {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	mostNum := MajorityElement(nums)

	log.Printf("数组%v最多的元素mostNum: %d", nums, mostNum)
}
