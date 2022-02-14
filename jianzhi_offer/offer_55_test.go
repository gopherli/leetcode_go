package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer55(b *testing.B) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	deep := MaxDepth(root)
	log.Println("二叉树的深度deep:", deep)
}
