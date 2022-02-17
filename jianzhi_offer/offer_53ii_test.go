package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer53ii(b *testing.B) {
	arr := []int{
		1, 2, 3,
	}
	num := MissingNumber(arr)
	fmt.Println("剑指 Offer 53 - II. 0～n-1中缺失的数字", num)
}
