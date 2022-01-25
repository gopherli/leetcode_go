package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer28(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := IsSymmetric(tree)

	log.Printf("是否是对称树：%v", result)
}
