package jianzhi_offer

// https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/

func MaxDepth(root *TreeNode) int {
	// 空树
	if root == nil {
		return 0
	}

	// 后序遍历：左子树，右子树dfs后的深度（最大）+1（根节点为第一层）
	return CompareInt(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

// 比较大小
func CompareInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
