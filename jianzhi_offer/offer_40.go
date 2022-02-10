package jianzhi_offer

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

func GetLeastNumbers(arr []int, k int) []int {
	// 判空处理
	if len(arr) == 0 {
		return []int{}
	}

	// 冒泡排序
	// sortArr := BubbleSort(arr)

	// 快速排序
	sortArr := QuickSort(arr, k)

	// 取排序数组前k个为新数组
	return sortArr[:k]
}

// 冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

// 快速排序
func QuickSort(arr []int, k int) []int {
	return _quickSort(arr, 0, len(arr)-1, k)
}

func _quickSort(arr []int, left, right, k int) []int {
	if left < right {
		// 分区基准
		partitionIndex := partition(arr, left, right)
		// 左分区排序
		_quickSort(arr, left, partitionIndex-1, k)
		// 右分区排序
		if partitionIndex < k {
			_quickSort(arr, partitionIndex+1, right, k)
		}
	}
	return arr
}

// 分区操作
func partition(arr []int, left, right int) int {
	// 基准
	privot := left

	// 起始下标
	index := privot + 1

	// 分区排序
	for i := index; i <= right; i++ {
		if arr[i] < arr[privot] {
			// 比基准小，交换元素
			swap(arr, i, index)

			// 起始下标顺移
			index++
		}
	}

	// 基准归位到对应位置
	swap(arr, privot, index-1)
	return index - 1
}

// 交换元素
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
