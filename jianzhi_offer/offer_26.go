package jianzhi_offer

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

/*
思路
1. 节点匹配，继续寻找
2. 节点不匹配，匹配失败
3. A树先为空，B树不为空，匹配失败
4. A树不为空，B树为空，匹配成功
*/

func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	/*V1*/
	// if A == nil && B != nil {
	// 	return false
	// }
	// if A == nil && B == nil {
	// 	return false
	// }

	// if A != nil && B == nil {
	// 	return false
	// }

	// if IsSame(A, B) {
	// 	return true
	// }

	// return IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B)

	/*V2*/
	return (A != nil && B != nil) && (Recur(A, B) || IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B))
}

func Recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return Recur(A.Left, B.Left) && Recur(A.Right, B.Right)
}

func IsSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A.Val == B.Val {
		return IsSame(A.Left, B.Left) && IsSame(A.Right, B.Right)
	}
	return false
}
