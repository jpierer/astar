package astar

import "errors"

// List represents a list of nodes
type List struct {
	nodes []Node
}

// NewList creates a new list
func NewList() *List {
	return &List{}
}

// Add a node to the list
func (l *List) Add(node Node) {
	l.nodes = append(l.nodes, node)
}

// Fill adds a slice of nodes to the list
func (l *List) Fill(nodes []Node) {
	l.nodes = append(l.nodes, nodes...)
}

// All returns the full list of nodes
func (l *List) All() []Node {
	return l.nodes
}

// Remove a node from the list
// if the node is not found we do nothing
func (l *List) Remove(removeNode Node) {
	index := l.GetIndex(removeNode)
	if index >= 0 {
		l.nodes = append(l.nodes[:index], l.nodes[index+1:]...)
	}
}

// GetIndex returns the index of the node in the list
// if the node is not found the return value is -1
func (l *List) GetIndex(searchNode Node) int {
	for index, node := range l.nodes {
		if node.X == searchNode.X && node.Y == searchNode.Y {
			return index
		}
	}
	return -1
}

// Contains check if a node is in the list
func (l *List) Contains(searchNode Node) bool {
	return l.GetIndex(searchNode) >= 0
}

// IsEmpty returns if the nodes list has nodes or not
func (l *List) IsEmpty() bool {
	return len(l.nodes) == 0
}

// Clear removes all nodes from the list
func (l *List) Clear() {
	l.nodes = []Node{}
}

// GetIndexOfMinF returns the index of the nodes list
// with the smallest node.F value
//
// if no node is found it returns -1
func (l *List) GetIndexOfMinF() int {
	lastNode := Node{}
	lastNodeIndex := -1
	for index, node := range l.nodes {
		if lastNodeIndex == -1 || node.F < lastNode.F {
			lastNode = node
			lastNodeIndex = index
		}
	}
	return lastNodeIndex
}

// GetMinF returns the node with the smallest node.F value
func (l *List) GetMinF() (Node, error) {
	minFIndex := l.GetIndexOfMinF()
	if minFIndex == -1 {
		return Node{}, errors.New("no node found")
	}
	return l.nodes[minFIndex], nil
}
