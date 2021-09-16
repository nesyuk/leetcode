package sorting

import (
	"testing"
	"reflect"
)

func TestQuickSort(t *testing.T) {
	tests := []struct{data []int; want []int}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3,2}, []int{2, 3}},
		{[]int{3,2,1}, []int{1,2,3}},
		{[]int{3,2,1,1,2,3}, []int{1,1,2,2,3,3}},
		{[]int{1, 10000000, 10, 100}, []int{1, 10, 100, 10000000}},
	}
	for _, test := range tests {
		QuickSort(test.data)
		if !reflect.DeepEqual(test.data, test.want) {
			t.Fatalf("expected: %v, got: %v", test.want, test.data)
		}
	}
}
