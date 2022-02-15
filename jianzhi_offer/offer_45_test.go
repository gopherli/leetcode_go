package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer45(b *testing.B) {
	arr := []int{3, 30, 34, 5, 9}
	min := MinNumber(arr)
	log.Println("把数组排成最小的数", min)
}
