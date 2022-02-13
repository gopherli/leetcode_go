package jianzhi_offer

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

func LevelOrder(root *TreeNode) []int {
	// 判空处理
	if root == nil {
		return []int{}
	}

	// 层次遍历结果
	arr := make([]int, 0)
	// 队列
	queue := make([]*TreeNode, 0)
	// 初始化队列
	queue = append(queue, root)

	// 队列不为空时，执行
	for len(queue) > 0 {
		// 读取队列第一个节点
		node := queue[0]
		arr = append(arr, node.Val)
		// 出队列
		queue = queue[1:]
		// 添加左子树
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		// 添加右子树
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return arr
}
