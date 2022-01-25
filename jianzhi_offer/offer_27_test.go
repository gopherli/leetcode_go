package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer27(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	log.Printf("原始树：%v", tree)

	mTree := mirrorTree(tree)
	log.Printf("镜像树：%v", mTree)
}
