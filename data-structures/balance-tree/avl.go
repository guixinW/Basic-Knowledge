package balance_tree

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
)

type AvlTree[T constraints.Ordered] struct {
	rootNode *TreeNode[T]
}

func NewAvlTree[T constraints.Ordered](rootVal T) *AvlTree[T] {
	newTree := new(AvlTree[T])
	newTree.rootNode = new(TreeNode[T])
	newTree.rootNode.Val = rootVal
	return newTree
}

// Insert 增加Val为insertVal的节点
func (t *AvlTree[T]) Insert(insertVal T) error {
	if t.rootNode == nil {
		return errors.New("rootNode is nil")
	}
	moveNode := t.rootNode
	parentOfMoveNode := moveNode
	for moveNode != nil {
		if insertVal > moveNode.Val {
			parentOfMoveNode = moveNode
			moveNode = moveNode.Right
		} else if insertVal < moveNode.Val {
			parentOfMoveNode = moveNode
			moveNode = moveNode.Left
		} else {
			break
		}
	}
	if moveNode != nil {
		return errors.New("move node already exists")
	}
	if parentOfMoveNode.Val < insertVal {
		parentOfMoveNode.Right = new(TreeNode[T])
		parentOfMoveNode.Right.Val = insertVal
	} else if parentOfMoveNode.Val > insertVal {
		parentOfMoveNode.Left = new(TreeNode[T])
		parentOfMoveNode.Left.Val = insertVal
	}
	return nil

}

// Find 查询Val为findVal的节点
func (t *AvlTree[T]) Find(findVal T) *TreeNode[T] {
	moveNode := t.rootNode
	for moveNode != nil {
		if findVal > moveNode.Val {
			moveNode = moveNode.Right
		} else if findVal < moveNode.Val {
			moveNode = moveNode.Left
		} else {
			break
		}
	}
	return moveNode
}

func (t *AvlTree[T]) Height() int {
	return 0
}

func (t *AvlTree[T]) Root() *TreeNode[T] {
	return t.rootNode
}

func (t *AvlTree[T]) PreOrderPrint() {
	var dfs func(node *TreeNode[T])
	dfs = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		fmt.Printf("%v ", node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	fmt.Println()
	dfs(t.rootNode)
}

func (t *AvlTree[T]) InOrderPrint() {
	var dfs func(node *TreeNode[T])
	dfs = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		dfs(node.Left)
		fmt.Printf("%v ", node.Val)
		dfs(node.Right)
	}
	dfs(t.rootNode)
}

func (t *AvlTree[T]) PostOrderPrint() {
	var dfs func(node *TreeNode[T])
	dfs = func(node *TreeNode[T]) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		fmt.Printf("%v ", node.Val)
	}
	dfs(t.rootNode)
}

// Delete 删除Val为deleteVal的节点
//func (t *AvlTree[T]) Delete(deleteVal *T) error {
//	return nil
//}

// Modify 修改将Val为modifyVal的节点修改为replaceVal的节点
//func (t *AvlTree[T]) Modify(modifyVal *T, replaceVal *T) error {
//	return nil
//}
