package matrix

import (
	"reflect"
	"testing"
	)

func TestUpdateMatrixDP(t *testing.T) {
	expectedDistances := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}
	distances := UpdateMatrixDP([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}})
	if !reflect.DeepEqual(expectedDistances, distances) {
		t.Fatalf("expected: %v, got: %v",expectedDistances,  distances)	
	}
	expectedDistances = [][]int{{0,1,0,1,2},{1,1,0,0,1},{0,0,0,1,0},{1,0,1,1,1},{1,0,0,0,1}}
	distances = UpdateMatrixDP([][]int{{0,1,0,1,1},{1,1,0,0,1},{0,0,0,1,0},{1,0,1,1,1},{1,0,0,0,1}})
	if !reflect.DeepEqual(expectedDistances, distances) {
		t.Fatalf("expected: %v, got: %v",expectedDistances,  distances)	
	}
}

func TestUpdateMatrixBST(t *testing.T) {
	expectedDistances := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}
	distances := UpdateMatrixBST([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}})
	if !reflect.DeepEqual(expectedDistances, distances) {
		t.Fatalf("expected: %v, got: %v",expectedDistances,  distances)	
	}
}
