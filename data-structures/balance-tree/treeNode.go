package balance_tree

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	Val    T
	Left   *TreeNode[T]
	Right  *TreeNode[T]
	Height int
}

func (node *TreeNode[T]) updateHeight() {
	var dfs func(node *TreeNode[T]) int
	dfs = func(node *TreeNode[T]) int {
		if node == nil {
			return 0
		}
		leftHeight := 1 + dfs(node.Left)
		rightHeight := 1 + dfs(node.Right)
		return max(leftHeight, rightHeight)
	}
	node.Height = dfs(node)
}

func (node *TreeNode[T]) GetBalanceFactor() int {
	return node.Left.Height - node.Right.Height
}

func (node *TreeNode[T]) GetHeight() int {
	return node.Height
}

func (node *TreeNode[T]) LeftRotate() *TreeNode[T] {
	t := node.Left
	node.Left = t.Right
	t.Right = node
	return t
}
