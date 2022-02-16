package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer46(b *testing.B) {
	num := 12258
	r := TranslateNum(num)
	fmt.Println("把数字翻译成字符串:", r)
}
