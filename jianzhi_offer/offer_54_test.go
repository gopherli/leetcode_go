package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer54(b *testing.B) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}},
	}

	// k := 1
	// k := 2
	k := 3
	// k := 4
	// k := 5

	val := KthLargest(root, k)

	log.Println("剑指 Offer 54. 二叉搜索树的第k大节点值:", val)
}
