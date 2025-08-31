package skipList

import (
	"golang.org/x/exp/rand"
)

type Skiplist struct {
	head  *Node
	level int
}

type Node struct {
	val     int
	forward []*Node
}

const MaxLevel = 32
const Probability = 0.25

func Constructor() Skiplist {
	ret := Skiplist{
		&Node{val: -1, forward: make([]*Node, MaxLevel)},
		0,
	}
	return ret
}

func (*Skiplist) randomLevel() int {
	i := 1
	for i <= MaxLevel && rand.Float64() < Probability {
		i++
	}
	return i
}

func (sl *Skiplist) Add(num int) {
	update := make([]*Node, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = sl.head
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	newLevel := sl.randomLevel()
	sl.level = max(newLevel, sl.level)
	newNode := &Node{num, make([]*Node, newLevel)}
	for i := 0; i < newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func (sl *Skiplist) Erase(num int) bool {
	cur := sl.head
	update := make([]*Node, MaxLevel)
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur == nil || cur.val != num {
		return false
	}
	for i := 0; i < sl.level && update[i].forward[i] == cur; i++ {
		update[i].forward[i] = cur.forward[i]
	}
	for sl.level > 0 && sl.head.forward[sl.level-1] == nil {
		sl.level--
	}
	return true
}

func (sl *Skiplist) Search(target int) bool {
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].val < target {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	if cur == nil || cur.val != target {
		return false
	}
	return true
}
