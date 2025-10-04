package heap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	newHeap := &Heap{}
	heap.Init(newHeap)
	heap.Push(newHeap, 1)
	heap.Push(newHeap, -3)
	heap.Push(newHeap, 5)
	heap.Push(newHeap, 2)
	heap.Push(newHeap, -1)
	if heap.Pop(newHeap).(int) != -3 {
		t.Error("Heap Pop failed")
	}
}
