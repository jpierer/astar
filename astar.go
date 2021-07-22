package astar

import (
	"fmt"
	"math"
)

// Node represents a simple node
type Node struct {
	F, G, H int
	X, Y    int
	Parent  *Node
}

type Grid struct {
	With, Height int
}

// String returns formatted values of the node
func (n Node) String() string {
	return fmt.Sprintf("Node [X:%d Y:%d F:%d G:%d H:%d]", n.X, n.Y, n.F, n.G, n.H)
}

// H caluclates the absolute distance between
// nodeA and nodeB calculates by the manhattan distance
func H(nodeA Node, nodeB Node) int {
	absX := math.Abs(float64(nodeA.X - nodeB.X))
	absY := math.Abs(float64(nodeA.Y - nodeB.Y))
	return int(absX + absY)
}

// 		nachbar knoten berechnen
// die berechnung ob der node später up, down, left, right ist, muss ich in der snake nochmal selbst programmieren
func GetNeighborNodes(node Node) []Node {
	var neighborNodes []Node

	upNode := Node{X: node.X, Y: node.Y + 1}
	if IsAccessible(upNode) {
		neighborNodes = append(neighborNodes, upNode)
	}

	downNode := Node{X: node.X, Y: node.Y - 1}
	if IsAccessible(downNode) {
		neighborNodes = append(neighborNodes, downNode)
	}

	leftNode := Node{X: node.X - 1, Y: node.Y}
	if IsAccessible(leftNode) {
		neighborNodes = append(neighborNodes, leftNode)
	}

	rightNode := Node{X: node.X + 1, Y: node.Y}
	if IsAccessible(rightNode) {
		neighborNodes = append(neighborNodes, rightNode)
	}

	// calculateNode(node, parentNode)?

	return neighborNodes
}

// IsAccessible checks if the node has valid coordinates
// and is not an obstacle
func IsAccessible(node Node) bool {
	// todo replace 7 with gridWidth and gridHeight

	// if node is out of bound
	if node.X < 0 || node.Y < 0 || node.X > 7-1 || node.Y > 7-1 {
		return false
	}

	// todo add more checks here
	// is node is in list of .... snake bodys etc...
	// 		ggf. über eine interface abstraction?

	return true
}

//func FindPath() ([]Node, error) {
//
//	return []Node{}, errors.New("No path found")
//}
