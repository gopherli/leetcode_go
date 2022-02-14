package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer57(b *testing.B) {
	arr := []int{14, 15, 16, 22, 53, 60}
	target := 76
	res := TwoSum(arr, target)
	log.Println("和为s的两个数字res：", res)
}
