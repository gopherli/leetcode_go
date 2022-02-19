package jianzhi_offer

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/

func KthLargest(root *TreeNode, k int) int {
	var val int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			val = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return val
}
