package jianzhi_offer

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/

func LevelOrderIII(root *TreeNode) [][]int {
	// 判空处理
	if root == nil {
		return nil
	}

	// 二维数组，返回结果
	arr := make([][]int, 0)
	// 队列，BFS
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// true偶数层，false奇数层
	judge := false

	// BFS
	for len(queue) > 0 {
		// 一维数组，接收当层结果
		curLevel := make([]int, 0)
		// 队列大小
		size := len(queue)

		// 遍历队列，获取当层结果
		for i := 0; i < size; i++ {
			// 读取队列元素
			node := queue[0]

			// 一维数组接收
			// if !judge {
			// 	// 奇数，添加到尾部
			// 	curLevel = append(curLevel, node.Val)
			// } else {
			// 	// 偶数，添加到头部
			// 	head := []int{node.Val}
			// 	head = append(head, curLevel...)
			// 	curLevel = head
			// }
			curLevel = append(curLevel, node.Val)

			// 已读元素出队
			queue = queue[1:]

			// 如果左子树不为空，入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// 如果右子树不为空，入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 二维数组，接收一维数组
		if !judge {
			// 奇数层
			arr = append(arr, curLevel)
		} else {
			// 偶数层
			// 数组反转
			reverseArr(curLevel)
			arr = append(arr, curLevel)
		}

		judge = !judge
	}

	return arr
}

// 数组反转
func reverseArr(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
