package dijkstra

import (
	"testing"
)

var matrix = [][]int{
	{0, 9, 0, 4, 0, 0},
	{9, 0, 8, 0, 1, 0},
	{0, 8, 0, 0, 3, 2},
	{4, 0, 0, 0, 0, 8},
	{0, 1, 3, 0, 0, 1},
	{0, 0, 2, 8, 1, 0},
}

func TestShortPathBetweenSourceAndDestination(t *testing.T) {
	shortPathBetweenSourceAndDestination(matrix, 0, 5)
}
