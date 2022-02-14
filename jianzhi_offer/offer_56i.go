package jianzhi_offer

import "fmt"

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

func SingleNumbers(nums []int) []int {
	// 1. 快排遍历
	// return quickSortFor56I(nums)

	// 2. 异或位运算
	ret := 0
	for _, v := range nums {
		ret ^= v
	}

	div := 1
	for div&ret == 0 {
		div <<= 1
	}

	a, b := 0, 0
	for _, v := range nums {
		if div&v == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

// 快排遍历
func QuickSortFor56I(nums []int) []int {
	// 判空处理
	if len(nums) == 0 {
		return nil
	}

	// 结果
	res := make([]int, 0)

	// 快排,O(logN)
	quickSort56I(nums, 0, len(nums)-1)
	fmt.Println("quick_nums:", nums)

	// 判断头
	if nums[0] < nums[1] {
		res = append(res, nums[0])
	}
	// 判断尾
	if nums[len(nums)-1] > nums[len(nums)-2] {
		res = append(res, nums[len(nums)-1])
	}

	// 遍历
	for j := 1; j < len(nums)-1; j++ {
		if nums[j] > nums[j-1] && nums[j] < nums[j+1] {
			res = append(res, nums[j])
		}
	}

	return res
}

// 快排
func quickSort56I(arr []int, left, right int) {
	_quickSort56I(arr, left, right)
}

// 分区
func _quickSort56I(arr []int, left, right int) {
	// 递归条件
	if left < right {
		// 分区基准
		partitionIndex := partition56I(arr, left, right)
		// 左分区排序
		_quickSort56I(arr, left, partitionIndex-1)
		// 右分区排序
		_quickSort56I(arr, partitionIndex+1, right)
	}
}

// 排序
func partition56I(arr []int, left, right int) int {
	// 基准值
	privot := left
	// 起始下标
	index := left + 1
	// 排序,升序
	for i := index; i <= right; i++ {
		// 左边值大于右边值
		if arr[privot] > arr[i] {
			// 交换
			swap56I(arr, i, index)
			// 下标右移
			index++
		}
	}
	// 基准值归位
	swap56I(arr, privot, index-1)

	// 返回分区值
	return index - 1
}

// 交换
func swap56I(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
