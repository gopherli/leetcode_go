package jianzhi_offer

// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

func MajorityElement(nums []int) int {
	// 数组不为空，无需判空处理

	// 方法一：哈希
	mostNum := has39(nums)

	// 方法二：快速排序取中位数
	// QuickSort39(nums, 0, len(nums)-1)
	// mostNum := nums[len(nums)/2]

	return mostNum
}

// 快排后取中位数
func QuickSort39(arr []int, left, right int) []int {
	return _quickSort39(arr, left, right)
}

// 分区操作
func _quickSort39(arr []int, left, right int) []int {
	if left < right {
		// 分区基准
		partitionIndex := partition39(arr, left, right)
		// 左分区排序
		_quickSort39(arr, left, partitionIndex-1)
		// 右分区排序
		_quickSort39(arr, partitionIndex+1, right)
	}
	return arr
}

// 分区排序
func partition39(arr []int, left, right int) int {
	// 基准值
	vioit := left
	// 起始下标
	index := left + 1

	// 遍历排序
	for i := index; i < len(arr); i++ {
		// 当前小于基准值，交换位置
		if arr[i] < arr[vioit] {
			swap(arr, i, index)
			index++
		}
	}

	// 基准值归位
	swap(arr, vioit, index-1)

	return index - 1
}

func has39(nums []int) int {
	// 数组长度
	n := len(nums)

	// 哈希结构记录每个元素出现的次数
	// 数组某个元素存在一半以上，长度为n/2
	has := make(map[int]int, n/2)

	// 遍历数组，has记录每个元素出现的次数
	for _, v := range nums {
		has[v]++
	}

	// 遍历哈希，找出出现次数最多的元素
	mostNum := 0   // 最大的数字
	mostCount := 0 // 最大的次数
	for number, count := range has {
		if count > mostCount {
			mostCount = count
			mostNum = number
		}
	}

	return mostNum
}
