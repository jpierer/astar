package astar

import "fmt"

// Node represents a simple node
// X and Y represents the nodes coordinates on the grid
//
// IMPORTANT: The grid coordinates starts on the "bottom left" -> X:0 / Y:0
//
// With the Weighting value you can set the nodes heavy grade
// so a node with mud or water are heavier as gras or street
type Node struct {
	f, g, h   int
	X, Y      int
	Weighting int
	parent    *Node
}

// String returns formatted values of the node
func (n Node) String() string {
	return fmt.Sprintf("Node [X:%d Y:%d F:%d G:%d H:%d]", n.X, n.Y, n.f, n.g, n.h)
}
