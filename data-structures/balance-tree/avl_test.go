package balance_tree

import (
	"fmt"
	"testing"
)

func TestPreOrderPrint(t *testing.T) {
	newTree := NewAvlTree(10)
	newTree.Insert(10)
	newTree.Insert(9)
	newTree.Insert(8)
	newTree.PreOrderPrint()
	newTree.rootNode = newTree.Root().LeftRotate()
	newTree.PreOrderPrint()
	newTree.Insert(7)
	newTree.Insert(6)
	newTree.PreOrderPrint()
	root8 := newTree.Find(8)
	root8 = root8.LeftRotate()
	newTree.PreOrderPrint()
}

func TestNodeHeight(t *testing.T) {
	newTree := NewAvlTree(10)
	newTree.Insert(20)
	newTree.Insert(9)
	newTree.Insert(30)
	newTree.Insert(40)
	newTree.Insert(50)
	newTree.rootNode.updateHeight()
	newTree.PreOrderPrint()
	fmt.Println(newTree.Root().Height)
}
