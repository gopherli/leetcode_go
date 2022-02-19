package jianzhi_offer

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/

func LowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// 保证p,q的大小顺序,p>q
	if p.Val < q.Val {
		p, q = q, p
	}

	// 寻找公共节点
	for root != nil {
		if root.Val > p.Val { // 左子树
			root = root.Left
		} else if root.Val < q.Val { // 右子树
			root = root.Right
		} else {
			break // root等于p,q再往下走说明不满足条件啦s
		}
	}

	return root
}
