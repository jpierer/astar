package astar

import (
	"testing"
)

func TestAstar_H(t *testing.T) {
	nodeA := Node{X: 0, Y: 0}
	nodeB := Node{X: 1, Y: 0}
	nodeC := Node{X: 1, Y: 1}
	nodeD := Node{X: 2, Y: 5}

	// invalid grid
	_, err := New(Config{GridWidth: 1, GridHeight: 1})
	if err == nil {
		t.Fatal("there should be a grid min error", err)
	}

	// setup a 6x6 grid
	a, err := New(Config{GridWidth: 6, GridHeight: 6})
	if err != nil {
		t.Fatal("there should be no error", err)
	}

	if a.H(nodeA, nodeB) != 1 {
		t.Error("should be 1")
	}
	if a.H(nodeA, nodeA) != 0 {
		t.Error("should be 0")
	}
	if a.H(nodeA, nodeC) != 2 {
		t.Error("should be 2")
	}
	if a.H(nodeB, nodeD) != 6 {
		t.Error("should be 6")
	}
}

func TestGetNeighborNodes(t *testing.T) {

	// setup a 4x4 grid
	a, err := New(Config{GridWidth: 4, GridHeight: 4})
	if err != nil {
		t.Fatal("there should be no error", err)
	}
	wantList := NewList()
	invalidList := NewList()

	defer func() {
		wantList.Clear()
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
		{X: 9, Y: 9},
	}

	wantList.Fill(validNeighbors)
	invalidList.Fill(invalidNeighbors)

	for index, neighbor := range a.GetNeighborNodes(node) {
		if wantList.Contains(neighbor) {
			wantList.Remove(neighbor)
		}
		if invalidList.Contains(neighbor) {
			t.Error("this neighbor node should be invalid: ", node)
		}
		if index > len(validNeighbors) {
			t.Error("more neighbors found as expected")
		}
	}

	if !wantList.IsEmpty() {
		t.Error("not all expected neighbors found: ", wantList.All())
	}

}

func TestGetNeighborNodesWithInvalidNodes(t *testing.T) {

	wantList := NewList()
	obstacleList := NewList()
	defer func() {
		wantList.Clear()
		obstacleList.Clear()
	}()

	node := Node{X: 2, Y: 2}

	validNeighbors := []Node{
		{X: 2, Y: 3}, // up
		{X: 3, Y: 2}, // right
	}

	obstacleNodes := []Node{
		{X: 2, Y: 1}, // down is not accessible
		{X: 1, Y: 2}, // left is not accessible
	}

	wantList.Fill(validNeighbors)
	obstacleList.Fill(obstacleNodes)

	// setup a 4x4 grid
	a, err := New(Config{GridWidth: 4, GridHeight: 4, InvalidNodes: obstacleNodes})
	if err != nil {
		t.Fatal("there should be no error", err)
	}

	for index, neighbor := range a.GetNeighborNodes(node) {
		if wantList.Contains(neighbor) {
			wantList.Remove(neighbor)
		}
		if obstacleList.Contains(neighbor) {
			t.Error("this neighbor node should be invalid: ", node)
		}
		if index > len(validNeighbors) {
			t.Error("more neighbors found as expected")
		}
	}

	if !wantList.IsEmpty() {
		t.Error("not all expected neighbors found: ", wantList.All())
	}

}

func TestAstar_FindPathA(t *testing.T) {

	// [ ] [ ] [ ] [ ] [ ]   S: StartNode
	// [ ] [O] [ ] [E] [ ]   E: EndNode
	// [ ] [ ] [O] [P] [ ]   O: ObstacleNode
	// [ ] [S] [P] [P] [ ]   P: Valid Path
	// [ ] [ ] [ ] [ ] [ ]

	startNode := Node{X: 1, Y: 1}
	endNode := Node{X: 3, Y: 3}
	obstacleNodes := []Node{
		{X: 1, Y: 3},
		{X: 2, Y: 2},
	}

	pathNodesToFind := []Node{
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 3, Y: 2},
		{X: 3, Y: 3},
	}

	pathList := NewList()

	defer func() {
		pathList.Clear()
	}()

	pathList.Fill(pathNodesToFind)

	// setup a 5x5 grid
	a, err := New(Config{GridWidth: 5, GridHeight: 5, InvalidNodes: obstacleNodes})
	if err != nil {
		t.Fatal("there should be no error", err)
	}
	foundPath, err := a.FindPath(startNode, endNode)
	if err != nil {
		t.Error("there should be a path", err)
	}

	for index, pathNode := range foundPath {
		if pathList.Contains(pathNode) {
			pathList.Remove(pathNode)
		}
		if index > len(pathNodesToFind) {
			t.Error("more path nodes found as expected")
		}
	}

	if !pathList.IsEmpty() {
		t.Error("not all expected path nodes found: ", pathList.All())
	}
}

func TestAstar_FindPathB(t *testing.T) {

	// [ ] [ ] [ ] [ ] [ ]   S: StartNode
	// [ ] [E] [P] [ ] [ ]   E: EndNode
	// [ ] [O] [P] [O] [O]   O: ObstacleNode
	// [ ] [O] [S] [ ] [ ]   P: Valid Path
	// [ ] [ ] [ ] [ ] [ ]

	startNode := Node{X: 2, Y: 1}
	endNode := Node{X: 1, Y: 3}
	obstacleNodes := []Node{
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
	}

	pathNodesToFind := []Node{
		{X: 2, Y: 2},
		{X: 2, Y: 3},
		{X: 1, Y: 3},
	}

	pathList := NewList()

	defer func() {
		pathList.Clear()
	}()

	pathList.Fill(pathNodesToFind)

	// setup a 5x5 grid
	a, err := New(Config{GridWidth: 5, GridHeight: 5, InvalidNodes: obstacleNodes})
	if err != nil {
		t.Fatal("there should be no error", err)
	}
	foundPath, err := a.FindPath(startNode, endNode)
	if err != nil {
		t.Error("there should be a path", err)
	}

	for index, pathNode := range foundPath {
		if pathList.Contains(pathNode) {
			pathList.Remove(pathNode)
		}
		if index > len(pathNodesToFind) {
			t.Error("more path nodes found as expected")
		}
	}

	if !pathList.IsEmpty() {
		t.Error("not all expected path nodes found: ", pathList.All())
	}
}

func TestAstar_FindPathC(t *testing.T) {

	// [ ] [ ] [ ] [ ] [ ]   S: StartNode
	// [ ] [ ] [ ] [ ] [ ]   E: EndNode
	// [ ] [ ] [ ] [ ] [ ]   O: ObstacleNode
	// [P] [E] [ ] [ ] [ ]   P: Valid Path
	// [S] [P] [ ] [ ] [ ]

	startNode := Node{X: 0, Y: 0}
	endNode := Node{X: 1, Y: 1}

	// way A
	pathA := []Node{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}

	// way B
	pathB := []Node{
		{X: 1, Y: 0},
		{X: 1, Y: 1},
	}

	pathAList := NewList()
	pathBList := NewList()

	defer func() {
		pathAList.Clear()
		pathBList.Clear()
	}()

	pathAList.Fill(pathA)
	pathBList.Fill(pathB)

	// setup a 5x5 grid
	a, err := New(Config{GridWidth: 5, GridHeight: 5})
	if err != nil {
		t.Fatal("there should be no error", err)
	}
	foundPath, err := a.FindPath(startNode, endNode)
	if err != nil {
		t.Error("there should be a path", err)
	}

	for index, pathNode := range foundPath {
		// check for way A
		if pathAList.Contains(pathNode) {
			pathAList.Remove(pathNode)
		}
		// check for way B
		if pathBList.Contains(pathNode) {
			pathBList.Remove(pathNode)
		}
		if index >= len(pathA)+len(pathB) {
			t.Error("we get 2 found path ways...")
		}
	}

	// one list must be empty here
	if !pathAList.IsEmpty() == !pathBList.IsEmpty() {
		t.Error("not all expected path nodes found!")
	}
}

func TestAstar_FindPathNoPath(t *testing.T) {

	// [ ] [ ] [ ] [ ] [E]   S: StartNode
	// [ ] [ ] [ ] [ ] [ ]   E: EndNode
	// [O] [O] [O] [O] [O]   O: ObstacleNode
	// [ ] [ ] [ ] [ ] [ ]   P: Valid Path
	// [S] [ ] [ ] [ ] [ ]

	startNode := Node{X: 0, Y: 0}
	endNode := Node{X: 4, Y: 4}
	obstacleNodes := []Node{
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
	}

	// setup a 5x5 grid
	a, err := New(Config{GridWidth: 5, GridHeight: 5, InvalidNodes: obstacleNodes})
	if err != nil {
		t.Fatal("there should be no error", err)
	}
	foundPath, err := a.FindPath(startNode, endNode)
	if err == nil {
		t.Error("there should be no path", err)
	}

	if len(foundPath) > 0 {
		t.Error("there should be no foundPath", foundPath)
	}

}
