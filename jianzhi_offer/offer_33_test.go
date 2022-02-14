package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer33(b *testing.B) {
	arr := []int{}
	bl := VerifyPostorder(arr)
	log.Println("叉搜索树的后序遍历序列:", bl)
}
