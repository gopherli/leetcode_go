package jianzhi_offer

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

func LevelOrderII(root *TreeNode) [][]int {
	// 判空处理
	if root == nil {
		return [][]int{}
	}

	// 二维数组接受遍历结果
	arrsTwo := make([][]int, 0)
	// 队列实现BFS
	queue := make([]*TreeNode, 0)
	// root先入队列
	queue = append(queue, root)

	// BFS
	for len(queue) > 0 {
		// 保存当前层次的元素
		curLevel := make([]int, 0)
		// 当前队列的元素
		size := len(queue)
		for i := 0; i < size; i++ {
			// 读取队列元素
			node := queue[0]
			curLevel = append(curLevel, node.Val)
			// 已读取元素，出队列
			queue = queue[1:]
			// 左子树不为空，入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 右子树不为空，入队列
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 当前层次元素数组添加至二维数组
		arrsTwo = append(arrsTwo, curLevel)
	}

	return arrsTwo
}
