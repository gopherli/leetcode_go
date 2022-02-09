package jianzhi_offer

import (
	"log"
	"testing"
)

//https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/

func TestOffer20(t *testing.T) {
	s := "e11"
	isNumber := IsNumber(s)
	log.Println("s是否是数值：", isNumber)
}
