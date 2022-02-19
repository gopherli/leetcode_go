package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer68(b *testing.B) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 13}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 33}},
	}
	p := root.Left
	q := root.Right
	c := LowestCommonAncestor(root, p, q)
	log.Println("剑指 Offer 68 - I. 二叉搜索树的最近公共祖先:", c.Val)
}
