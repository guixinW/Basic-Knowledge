package dijkstra

import (
	"container/heap"
	"math"
)

type Edges struct {
	weight int
	node   int
}

type PriorityQueue []Edges

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edges))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func shortPathBetweenSourceAndDestination(edges [][]int, source, dest int) int {
	nodeEdges := make(map[int][]Edges)
	distance := make([]int, len(edges))
	for i := 0; i < len(distance); i++ {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0
	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			if edges[i][j] != 0 {
				nodeEdges[i] = append(nodeEdges[i], Edges{edges[i][j], j})
				nodeEdges[j] = append(nodeEdges[j], Edges{edges[i][j], i})
			}
		}
	}
	visitedNode := make([]bool, len(edges))
	edgesQueue := make(PriorityQueue, 0)
	heap.Init(&edgesQueue)
	heap.Push(&edgesQueue, Edges{0, source})
	for len(edgesQueue) > 0 {
		shortestEdges := heap.Pop(&edgesQueue).(Edges)
		if visitedNode[shortestEdges.node] {
			continue
		}
		visitedNode[shortestEdges.node] = true
		if shortestEdges.node == dest {
			break
		}
		nowNodeEdges := nodeEdges[shortestEdges.node]
		for i := 0; i < len(nowNodeEdges); i++ {
			if visitedNode[nowNodeEdges[i].node] == false {
				if distance[nowNodeEdges[i].node] > distance[shortestEdges.node]+nowNodeEdges[i].weight {
					distance[nowNodeEdges[i].node] = distance[shortestEdges.node] + nowNodeEdges[i].weight
					heap.Push(&edgesQueue, Edges{distance[nowNodeEdges[i].node], nowNodeEdges[i].node})
				}
			}
		}
	}
	return distance[dest]
}
