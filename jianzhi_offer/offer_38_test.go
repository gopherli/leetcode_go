package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer38(t *testing.T) {
	string := "abbs"

	ans := Permutation(string)

	log.Printf("字符串%s的不重复全排列结果ans:%v", string, ans)
}
