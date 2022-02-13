package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer32(b *testing.B) {
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
	arr := LevelOrder(root)
	log.Printf("二叉树层次遍历arr：%v", arr)
}
