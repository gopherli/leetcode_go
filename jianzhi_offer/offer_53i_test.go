package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer53i(b *testing.B) {
	target := 0
	arr := []int{5, 7, 7, 8, 8, 10}
	count := Search(arr, target)
	fmt.Println("剑指 Offer 53 - I. 在排序数组中查找数字 I出现的次数：", count)
}
