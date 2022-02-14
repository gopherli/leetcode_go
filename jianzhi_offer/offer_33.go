package jianzhi_offer

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

func VerifyPostorder(postorder []int) bool {
	// 二叉搜索树特点
	// 左子树节点值 < 根节点值
	// 根节点值 < 右子树节点值

	// 后序遍历特点
	// 左子树->右子树->根节点

	// 二叉搜索树后序遍历特点
	// [左子树全部节点<右子树全部节点>根节点]

	// 解题思路：
	// 已知：根节点，postorder[len(postorder)-1]
	// 可求得：左子树区间[0,m-1]，右子树区间[m,len(postorder)-2]
	// 转化：左子树、右子树可分别看成完整二叉搜索树->递归即可
	// 后序遍历有对应的二叉搜索树的条件为：每个元素都被判断，然后累计判断数count == len(postorder)-1
	return _recur(postorder, 0, len(postorder)-1)
}

// 递归
func _recur(postorder []int, i, j int) bool {
	// 特殊处理:数组只有一个元素或者没有元素
	if i >= j {
		return true
	}

	// 记录判断count
	p := i

	// 根节点
	root := postorder[j]

	// 根节点>左子树
	for root > postorder[p] {
		p++
	}

	// 左右子树分届点
	m := p

	// 根节点<右子树
	for root < postorder[p] {
		p++
	}

	// 递归
	return p == j && _recur(postorder, i, m-1) && _recur(postorder, m, j-1)
}
