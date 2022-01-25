package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer26(t *testing.T) {
	ATree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	BTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	result := IsSubStructure(ATree, BTree)
	log.Fatalf("B树是A树的子树：%v", result)
}
