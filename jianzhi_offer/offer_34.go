package jianzhi_offer

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

func PathSum(root *TreeNode, target int) [][]int {
	// 路径
	path := make([]int, 0)

	// 结果
	res := make([][]int, 0)

	// 定义dfs
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, tar int) {
		// 如果节点为空，方法return
		if node == nil {
			return
		}

		// 目标值不断减去当前路径上节点值
		tar = tar - node.Val

		// 路径添加节点值
		path = append(path, node.Val)

		// defer(return后执行)去掉叶子节点，拼接其他叶子节点
		defer func() {
			// path变成了可变变量
			path = path[:len(path)-1]
		}()

		// 叶子节点：左子树、右子树为空
		// 满足路径值：tar = 0
		if node.Left == nil && node.Right == nil && tar == 0 {
			// res添加一个符合结果的path
			// 由于path是可变变量，需要path...
			res = append(res, append([]int(nil), path...))
			return
		}

		// 左子树
		dfs(node.Left, tar)

		// 右子树
		dfs(node.Right, tar)
	}

	// 执行dfs
	dfs(root, target)

	return res
}
