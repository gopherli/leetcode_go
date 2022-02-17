package jianzhi_offer

import (
	"fmt"
	"testing"
)

func BenchmarkOffer58i(b *testing.B) {
	str1 := "the sky is blue"
	str2 := "  hello world!  "
	str3 := "a good   example"

	rStr1 := ReverseWords(str1)
	rStr2 := ReverseWords(str2)
	rStr3 := ReverseWords(str3)

	fmt.Println("rStr1:", rStr1)
	fmt.Println("rStr2:", rStr2)
	fmt.Println("rStr3:", rStr3)

}
