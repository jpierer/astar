package astar

import "testing"

func TestList_Add(t *testing.T) {
	nodeA := Node{}
	nodeB := Node{}

	list := NewList()

	list.Add(nodeA)

	if len(list.nodes) != 1 {
		t.Error("should have a node")
	}

	list.Add(nodeB)

	if len(list.nodes) < 2 {
		t.Error("should have a 2 nodes")
	}
}

func TestList_Contains(t *testing.T) {
	nodeA := Node{X: 0, Y: 0}
	nodeB := Node{X: 0, Y: 1, Parent: &nodeA}

	list := NewList()

	list.Add(nodeB)
	list.Add(nodeA)

	if !list.Contains(nodeA) {
		t.Error("should have nodeA")
	}

	if !list.Contains(nodeA) {
		t.Error("should have nodeB")
	}
}

func TestList_Remove(t *testing.T) {
	nodeA := Node{X: 0, Y: 0}
	nodeB := Node{X: 0, Y: 1, Parent: &nodeA}

	list := NewList()

	list.Add(nodeB)
	list.Add(nodeA)

	list.Remove(nodeA)

	if list.IsEmpty() {
		t.Error("IsEmpty should be false")
	}

	if list.Contains(nodeA) {
		t.Error("nodeA should not exist")
	}

	if !list.Contains(nodeB) {
		t.Error("nodeB should still be there")
	}

	list.Remove(nodeB)
	list.Remove(nodeB) // try removing it twice
	if list.Contains(nodeB) {
		t.Error("nodeB should now removed")
	}

	if !list.IsEmpty() {
		t.Error("IsEmpty should be true")
	}
}

func TestList_Clear(t *testing.T) {
	nodeA := Node{X: 0, Y: 0}
	nodeB := Node{X: 0, Y: 1, Parent: &nodeA}

	list := NewList()

	list.Add(nodeB)
	list.Add(nodeA)

	list.Clear()

	if !list.IsEmpty() {
		t.Error("IsEmpty should be true")
	}
}

func TestList_GetIndexOfMinF(t *testing.T) {
	nodeA := Node{X: 0, Y: 0, F: 2}
	nodeB := Node{X: 0, Y: 1, F: 3, Parent: &nodeA}
	nodeC := Node{X: 0, Y: 1, F: 4, Parent: &nodeB}

	list := NewList()

	list.Add(nodeB)
	list.Add(nodeC)
	list.Add(nodeA)

	if list.GetIndexOfMinF() != 2 {
		t.Error("Should be index 2")
	}

	nodeD := Node{X: 0, Y: 1, F: 1, Parent: &nodeC}
	list.Add(nodeD)

	if list.GetIndexOfMinF() != 3 {
		t.Error("Should be index 3")
	}

	list.Clear()

	if list.GetIndexOfMinF() != -1 {
		t.Error("Should be index -1")
	}
}

func TestList_GetMinF(t *testing.T) {
	nodeA := Node{X: 1, Y: 0, F: 2}
	nodeB := Node{X: 2, Y: 2, F: 3, Parent: &nodeA}
	nodeC := Node{X: 2, Y: 3, F: 4, Parent: &nodeB}

	list := NewList()

	list.Add(nodeB)
	list.Add(nodeA)
	list.Add(nodeC)

	node, err := list.GetMinFNode()
	if err != nil {
		t.Error("There should be no error", err)
	}

	if node.X != nodeA.X || node.Y != nodeA.Y {
		t.Error("This is not the correct node we expect")
	}

	list.Clear()

	if _, err := list.GetMinFNode(); err == nil {
		t.Error("we should have an error here")
	}
}
