package jianzhi_offer

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/

func SpiralOrder(matrix [][]int) []int {
	// 空数组返回空
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 行列
	rows := len(matrix)
	columns := len(matrix[0])

	// 上下左右
	top, bottom, left, right := 0, rows-1, 0, columns-1

	// 结果数组
	order := make([]int, rows*columns)

	// 结果数组下标
	index := 0

	// 圈层遍历
	for left <= right && top <= bottom {
		//上(top,left) (top,right), 行不变列变
		for moveLeft := left; moveLeft <= right; moveLeft++ {
			order[index] = matrix[top][moveLeft]
			index++
		}

		//右(top+1,right) (bottom,right)，列不变行变
		for moveTop := top + 1; moveTop <= bottom; moveTop++ {
			order[index] = matrix[moveTop][right]
			index++
		}

		//下(bottom,right-1) (bottom,left)，行不变列变
		for moveRight := right - 1; moveRight >= left && bottom != top; moveRight-- {
			order[index] = matrix[bottom][moveRight]
			index++
		}

		//左(bottom-1,left) (top,left)，列不变行变
		for moveBottom := bottom - 1; moveBottom > top && left != right; moveBottom-- {
			order[index] = matrix[moveBottom][left]
			index++
		}

		// 遍历下一个圈层
		top++
		bottom--
		left++
		right--
	}

	return order
}
