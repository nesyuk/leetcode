package graph

import (
	"testing"
	"reflect"
)

func TestCheckNodesConnected(t *testing.T) {
	node3 := &Node{Val: 3, Neighbours: []*Node{}}
	node2 := &Node{Val: 2, Neighbours: []*Node{node3}}
	node1 := &Node{Val: 1, Neighbours: []*Node{node2}}
	node2.Neighbours = append(node2.Neighbours, node2)
	node3.Neighbours = append(node3.Neighbours, node2)
	if CheckNodesConnected(node1, node3, map[*Node]bool{node1: true}) != true {
		t.Fatalf("expected: %v, got: %v", true, false)
	}
	if CheckNodesConnected(node3, node1, map[*Node]bool{node1: true}) != false {
		t.Fatalf("expected: %v, got: %v", false, true)
	}
	
}

func TestClone(t *testing.T) {
	node3 := &Node{Val: 3, Neighbours: []*Node{}}
        node2 := &Node{Val: 2, Neighbours: []*Node{node3}}
        node1 := &Node{Val: 1, Neighbours: []*Node{node2}}
        node2.Neighbours = append(node2.Neighbours, node2)
        node3.Neighbours = append(node3.Neighbours, node2)

	node1Clone := Clone(node1)
	if !reflect.DeepEqual(node1, node1Clone) {
		t.Fatalf("expected: %v, got: %v", node1, node1Clone)
	}
}
