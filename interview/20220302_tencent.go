package interview

import "math"

// 1、大数相乘
func BigNum(bigOne, bigTwo int) int {
	// 结果
	result := 0

	// 整型的最大范围
	rangeIntMax := math.MaxInt

	// 默认整形int的话，输入的两个数是不会超出的，也不判断
	// 相乘不能超过整型的最大范围
	result = bigOne * bigTwo

	// 结果不能超出范围，即返回
	// 超出范围用某种方式来规避
	if result > rangeIntMax {
		return -1
	}

	return result
}

type List struct {
	Val  int
	Next *List
}

// 2、两个有序链表合并为一个有序链表
func TwoListMerge(rootOne, rootTwo *List) *List {
	// 1->2->2->3
	// 2->3->4->5

	// 引入空节点收，空间耗费大
	rootMerge := &List{}

	// 遍历条件：两个链表存在一个不为空可以继续执行
	for rootOne != nil && rootTwo != nil {
		// 如果存在一个链表为空，另外一个不为空的情况，直接返回
		if rootOne == nil || rootTwo == nil {

			break
		}

		// 判断值的大小，进行有序拼接
		// 节点一头节点<=节点二头节点
		//
		if rootOne.Val <= rootTwo.Val {
			// 保留当前节点
			curOne := rootOne
			rootMerge.Next = curOne
			curOne.Next = nil
			// rootOne瞬移到下一个
			rootOne = rootOne.Next
		} else {
			curTwo := rootTwo
			rootMerge.Next = curTwo
			curTwo.Next = nil
			rootTwo = rootTwo.Next
		}

	}

	return rootMerge
}

// 3、LRU实现
func LRU() {
	// LRU: 最近最少使用
	// 经常使用过的添加到头部
	// 不经常使用的自然会存在于尾部

	// 1. 链表+哈希+排序
	// 链表：从头到尾-->高频到低频
	// 哈希：记录节点的访问次数
	// 排序：保证链表从头到尾是高频到低频（排序时机，不能一次断链重组就进行，消耗太大）

	// 节点：node1，node2，node3，node4，node5

	// 初始化：空链表，维护上述规则

	// root := &List{}
	// node1先被使用
	// root = node1
	// has[node1]++ // 记录被使用过一次，每使用一次进行计数+1
	// 假设按顺序被使用过：node2，node3，node4，node5
	// root = node1->node2->node3->node4->node5
	// node4 used
	// 遍历root,找到node4，执行断链重组，将node4放到头部(会存在问题，一直没有使用的突然使用一次就放到头部)
	// 如此反复，链表头部的节点使用频率高，尾部使用频率会相对较低

	// 栈：后进先出
	// 队列：先进先出
	// 双队列1: 先进先出
	// 双队列2: 先进先出
	// count--
	//  1 2 3 5-> top
	//  j o j o
	// (1)f 1 3  w  j
	// (2)f 2 5  w  0
}
