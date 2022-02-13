package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer39(t *testing.T) {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	mostNum := MajorityElement(nums)

	log.Printf("数组%v最多的元素mostNum：%d", nums, mostNum)
}
