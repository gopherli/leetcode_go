package jianzhi_offer

import (
	"log"
	"testing"
)

func BenchmarkOffer32III(b *testing.B) {
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
	arrsTwo := LevelOrderIII(root)
	log.Printf("二叉树S遍历二维数组接受arrsTwo：%v", arrsTwo)
}
