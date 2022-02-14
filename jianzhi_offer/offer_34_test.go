package jianzhi_offer

import (
	"log"
	"testing"
)

func TestOffer34(t *testing.T) {
	head := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 6}},
	}
	//   5
	// 4   8
	//3 2 7  6
	res := PathSum(head, 11)
	log.Printf("二叉树中和为某一值的路径res：%v", res)
}
