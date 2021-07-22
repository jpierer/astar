package astar

import "fmt"

// Node represents a simple node
type Node struct {
	F, G, H int
	X, Y    int
	Parent  *Node
}

// String returns formatted values of the node
func (n Node) String() string {
	return fmt.Sprintf("Node [X:%d Y:%d F:%d G:%d H:%d]", n.X, n.Y, n.F, n.G, n.H)
}
