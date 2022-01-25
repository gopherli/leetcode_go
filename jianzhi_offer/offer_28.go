package jianzhi_offer

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/

func IsSymmetric(root *TreeNode) bool {
	return IsSymmetricSame(root, root)
}

func IsSymmetricSame(left *TreeNode, right *TreeNode) bool {
	// 如果为空，对称
	if left == nil && right == nil {
		return true
	}
	// 如果不为空，左右子树所有节点值相等
	if left != nil && right != nil && left.Val == right.Val {
		// 递归走完剩下子树节点
		return IsSymmetricSame(left.Left, right.Right) && IsSymmetricSame(left.Right, right.Left)
	}

	return false
}
