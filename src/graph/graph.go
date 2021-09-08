package graph

import (
	"reflect"
)

type Node struct {
	Val int
	Neighbours []*Node
}

func Clone(node *Node) *Node {
	if node == nil {
		return nil
	}
	return clone(node, &map[*Node]*Node{})
}

func clone(node *Node, seen *map[*Node]*Node) *Node {
	nodeCpy := &Node{Val: node.Val, Neighbours: []*Node{}}
	(*seen)[node] = nodeCpy
	for _, neighbour := range node.Neighbours {
		if neighbourCpy, exist := (*seen)[neighbour]; exist {
			nodeCpy.Neighbours = append(nodeCpy.Neighbours, neighbourCpy)
		} else {
			nodeCpy.Neighbours = append(nodeCpy.Neighbours, clone(neighbour, seen))
		}
	}
	return nodeCpy
} 

func CheckNodesConnected(source *Node, target *Node, visited map[*Node]bool) bool {
	if reflect.DeepEqual(source, target) {
		return true
	}
	for _, child := range source.Neighbours {
		if _, isVisited := visited[child]; !isVisited {
			visited[child] = true
			return CheckNodesConnected(child, target, visited)	
		}
	}
	return false
}
