package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer63(b *testing.B) {
	arr := []int{
		7, 1, 5, 3, 6, 4,
	}

	max := MaxProfit(arr)

	log.Println("剑指 Offer 63. 股票的最大利润:", max)
}
