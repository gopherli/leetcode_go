package jianzhi_offer

import (
	"log"
	"testing"
)

func Benchmark40_01(b *testing.B) {
	arr := []int{3, 2, 1}
	k := 2
	kArr := GetLeastNumbers01(arr, k)
	log.Println("K个数组：", kArr)
}
