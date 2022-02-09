package jianzhi_offer

// https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

var cacheNode map[*Node]*Node

func CopyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return DeepCopy(head)
}

func DeepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, has := cacheNode[node]; has {
		return n
	}

	newHead := &Node{Val: node.Val}
	cacheNode[node] = newHead
	newHead.Next = DeepCopy(node.Next)
	newHead.Random = DeepCopy(node.Random)

	return newHead
}
