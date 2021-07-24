# astar 

a* (a-star) pathfinding algorithm written in go

Install
-------
    go get github.com/jpierer/astar@main

Example
-------

```go
package main

import (
	"fmt"
	"github.com/jpierer/astar"
)

func main() {

	// Example 5x5 Grid
	//
	// [ ] [ ] [ ] [ ] [ ]   S: StartNode    (The node were you are)
	// [ ] [E] [P] [ ] [ ]   E: EndNode      (The destination point where you want to go)
	// [ ] [O] [P] [O] [O]   O: ObstacleNode (Some obstacles you cannot access)
	// [W] [O] [S] [ ] [ ]   P: Valid Path   (This is just a visualisation of the returned found path)
	// [W] [ ] [ ] [ ] [ ]   W: WeightedNode (Nodes like water, which are harder to enter)
	//
	// IMPORTANT: The grid coordinates starts on the "bottom left" -> X:0 / Y:0

	startNode := astar.Node{X: 2, Y: 1}
	endNode := astar.Node{X: 1, Y: 3}

	obstacleNodes := []astar.Node{
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
	}
	waterNodes := []astar.Node{
		{X: 0, Y: 0, Weighting: 20},
		{X: 0, Y: 1, Weighting: 20},
	}

	// set nodes to the config
	aConfig := astar.Config{
		GridWidth:     5,
		GridHeight:    5,
		InvalidNodes:  obstacleNodes,
		WeightedNodes: waterNodes,
	}

	// create the algo with defined config
	algo, err := astar.New(aConfig)
	if err != nil {
		fmt.Println("invalid astar config", err)
		return
	}

	// run it
	foundPath, err := algo.FindPath(startNode, endNode)
	if err != nil || len(foundPath) == 0 {
		fmt.Println("No path found ...")
		return
	}

	// the foundPath has now the way to the target

	// IMPORTANT:
	// the path is in the opposite way so the endpoint node is on index 0
	// you can avoid it by switching the startNode<>endNode parameter
	for _, node := range foundPath {
		fmt.Println(node)
	}

	// output:
	// Node [X:1 Y:3 F:1 G:1 H:0]
	// Node [X:2 Y:3 F:2 G:1 H:1]
	// Node [X:2 Y:2 F:3 G:1 H:2]

}

```
