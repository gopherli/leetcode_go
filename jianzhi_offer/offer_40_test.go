package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer40(t *testing.T) {
	arr := []int{3, 2, 1, 0}
	k := 2
	kArr := GetLeastNumbers(arr, k)

	log.Printf("剑指 Offer 40. 最小的%v个数:%v", k, kArr)
}
