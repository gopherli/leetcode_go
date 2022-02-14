package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer56I(b *testing.B) {
	arr := []int{1, 2, 10, 4, 1, 4, 3, 3}
	res := SingleNumbers(arr)
	log.Println("数组中数字出现的次数:", res)
}
