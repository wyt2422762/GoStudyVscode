package main

// 在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
// 如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
// 我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
// 只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode //x, y的父节点
	var xDepth, yDepth int         //x, y的深度
	var xFound, yFound bool        //x, y是否找到

	var f func(node, parent *TreeNode, depth int) //深度优先搜索找x, y
	f = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		}
		if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		// 如果两个节点都找到了，就可以提前退出遍历
		if xFound && yFound {
			return
		}

		f(node.Left, node, depth+1)

		// 如果两个节点都找到了，就可以提前退出遍历
		if xFound && yFound {
			return
		}

		f(node.Right, node, depth+1)
	}

	f(root, nil, 0)

	return xDepth == yDepth && xParent != yParent //判断x, y的深度是否一致，以及是否是同一个父节点
}
