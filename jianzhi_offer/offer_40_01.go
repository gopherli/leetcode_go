package jianzhi_offer

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

func GetLeastNumbers01(arr []int, k int) []int {
	// 判空处理
	if len(arr) == 0 {
		return []int{}
	}

	// 快排
	QuickSortGetLeastNumbers01(arr, 0, len(arr)-1)

	// k
	return arr[:k]
}

func QuickSortGetLeastNumbers01(arr []int, left, right int) {
	if left < right {
		privot := left
		index := left + 1
		for i := index; i <= right; i++ {
			if arr[privot] > arr[i] {
				arr[index], arr[i] = arr[i], arr[index]
				index++
			}
		}
		arr[privot], arr[index-1] = arr[index-1], arr[privot]
		partitionIndex := index - 1
		QuickSortGetLeastNumbers01(arr, left, partitionIndex-1)
		QuickSortGetLeastNumbers01(arr, partitionIndex+1, right)
	}
}
