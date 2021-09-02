package main

import (
	"fmt"
	"bstree"
)

func postorderTraversalStacks(root *bstree.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*bstree.TreeNode{root}
	seen := []*bstree.TreeNode{}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		seen = append(seen, current)
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}
	for i := len(seen)-1; i >= 0; i-- {
		result = append(result, seen[i].Val)
	}
	return result
}

func run() {
	fmt.Println("go!")
	bs := bstree.NewBSTree([]int{2, 1, 3})
	fmt.Println(postorderTraversalStacks(bs))

	bs = bstree.NewBSTree([]int{5, 6, 3, 1, 5, 7, 8, 9, 2})
	fmt.Println(postorderTraversalStacks(bs))
}

func main() {
	run()
}
