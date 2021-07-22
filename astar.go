package astar

import (
	"errors"
	"math"
)

// Config holds important settings
// to perform the calculation
//
// GridWidth and GridHeight are required and represents
// the size of the grid
//
// InvalidNodes can be used to add obstacles nodes etc.
type Config struct {
	GridWidth, GridHeight int
	InvalidNodes          []Node
}

type astar struct {
	config              Config
	openList, closeList List
}

// New creates a new astar instance
func New(config Config) (*astar, error) {
	if config.GridWidth < 2 || config.GridHeight < 2 {
		return nil, errors.New("GridWidth and GridHeight must be min 2")
	}
	a := &astar{config: config}
	return a.init(), nil
}

func (a *astar) init() *astar {
	// add invalidNodes directly to the closeList
	a.closeList.Fill(a.config.InvalidNodes)
	return a
}

// H caluclates the absolute distance between
// nodeA and nodeB calculates by the manhattan distance
func (a *astar) H(nodeA Node, nodeB Node) int {
	absX := math.Abs(float64(nodeA.X - nodeB.X))
	absY := math.Abs(float64(nodeA.Y - nodeB.Y))
	return int(absX + absY)
}

// GetNeighborNodes calculates the next neighbors of the given node
// if a neighbor node is not accessible the node will be ignored
func (a *astar) GetNeighborNodes(node Node) []Node {
	var neighborNodes []Node

	upNode := Node{X: node.X, Y: node.Y + 1}
	if a.IsAccessible(upNode) {
		neighborNodes = append(neighborNodes, upNode)
	}

	downNode := Node{X: node.X, Y: node.Y - 1}
	if a.IsAccessible(downNode) {
		neighborNodes = append(neighborNodes, downNode)
	}

	leftNode := Node{X: node.X - 1, Y: node.Y}
	if a.IsAccessible(leftNode) {
		neighborNodes = append(neighborNodes, leftNode)
	}

	rightNode := Node{X: node.X + 1, Y: node.Y}
	if a.IsAccessible(rightNode) {
		neighborNodes = append(neighborNodes, rightNode)
	}

	// calculateNode(node, parentNode)?

	return neighborNodes
}

// IsAccessible checks if the node is reachable in the grid
// and is not in the invalidNodes slice
func (a *astar) IsAccessible(node Node) bool {

	// if node is out of bound
	if node.X < 0 || node.Y < 0 || node.X > a.config.GridWidth-1 || node.Y > a.config.GridHeight-1 {
		return false
	}

	// check if the node is in the closeList
	// the predefined invalidNodes are also in this list
	if a.closeList.Contains(node) {
		return false
	}

	return true
}

func (a *astar) FindPath() ([]Node, error) {

	defer func() {
		a.openList.Clear()
		a.closeList.Clear()
	}()

	// todo add calculation

	return []Node{}, errors.New("No path found")
}
