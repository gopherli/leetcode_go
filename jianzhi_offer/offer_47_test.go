package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer47(b *testing.B) {
	mn := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	maxPath := MaxValue(mn)

	fmt.Println("剑指 Offer 47. 礼物的最大价值:", maxPath)
}
