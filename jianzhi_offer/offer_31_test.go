package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer31(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 1, 2}

	bl := ValidateStackSequences(pushed, popped)

	log.Println("popped是pushed的一个弹出序列：", bl)
}
