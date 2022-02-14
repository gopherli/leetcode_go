package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer56II(b *testing.B) {
	arr := []int{0, 0, 0, 1, 2, 2, 2, 3, 3, 3}
	num := SingleNumber56II(arr)
	log.Println("数组中数字出现的次数II:", num)
}
