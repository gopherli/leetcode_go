package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer44(b *testing.B) {
	n := 11
	num := FindNthDigit(n)
	fmt.Println("剑指 Offer 44. 数字序列中某一位的数字:", num)
}
