package astar

import (
	"testing"
)

func TestAstar_H(t *testing.T) {
	nodeA := Node{X: 0, Y: 0}
	nodeB := Node{X: 1, Y: 0}
	nodeC := Node{X: 1, Y: 1}
	nodeD := Node{X: 2, Y: 5}

	if H(nodeA, nodeB) != 1 {
		t.Fatal("should be 1")
	}
	if H(nodeA, nodeA) != 0 {
		t.Fatal("should be 0")
	}
	if H(nodeA, nodeC) != 2 {
		t.Fatal("should be 2")
	}
	if H(nodeB, nodeD) != 6 {
		t.Fatal("should be 6")
	}
}

func TestGetNeighborNodes(t *testing.T) {

	validList := NewList()
	invalidList := NewList()

	defer func() {
		validList.Clear()
		invalidList.Clear()
	}()

	node := Node{X: 2, Y: 2}

	validNeighbors := []Node{
		{X: 2, Y: 3}, // up
		{X: 2, Y: 1}, // down
		{X: 1, Y: 2}, // left
		{X: 3, Y: 2}, // right
	}

	invalidNeighbors := []Node{
		{X: 0, Y: 3},
		{X: 0, Y: 2},
		{X: 0, Y: 0},
		{X: 3, Y: 3},
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 2, Y: 0},
	}

	for _, neighbor := range validNeighbors {
		validList.Add(neighbor)
	}

	for _, neighbor := range invalidNeighbors {
		invalidList.Add(neighbor)
	}

	for _, neighbor := range GetNeighborNodes(node) {
		if validList.Contains(neighbor) {
			validList.Remove(neighbor)
		}
		if invalidList.Contains(neighbor) {
			t.Fatal("this neighbor node should be invalid: ", node)
		}
	}

	if !validList.IsEmpty() {
		t.Fatal("not all neighbors found: ", validList.All())
	}

}
