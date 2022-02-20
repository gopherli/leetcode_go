package jianzhi_offer

// https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Abs(Depth(root.Left), Depth(root.Right)) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

// 深度：Max(left,right)+1
func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(Depth(root.Left), Depth(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
