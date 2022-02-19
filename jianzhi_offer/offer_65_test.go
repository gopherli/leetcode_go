package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer65(b *testing.B) {
	x := 5
	y := 6
	sum := Add(x, y)
	log.Println("剑指 Offer 65. 不用加减乘除做加法:", sum)
	multi := Multi(x, y)
	log.Println("剑指 Offer 65. 不用加减乘除做乘法:", multi)

}
