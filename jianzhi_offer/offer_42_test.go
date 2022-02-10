package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer42(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	max := MaxSubArray(nums)

	log.Println("连续子数组的最大和max=", max)
}
