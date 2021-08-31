package main

import (
	"fmt"
	"bstree"
)

func preorderTraversalStack(root *bstree.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*bstree.TreeNode{root}
	for len(stack) != 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, last.Val)
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
	}
	return result
}

func run() {
	fmt.Println("go!")
	bs := bstree.NewTree([]int{2, 1, 3})
	fmt.Println(preorderTraversalStack(bs))

	bs = bstree.NewTree([]int{5, 6, 3, 1, 5, 7, 8, 9, 2})
	fmt.Println(preorderTraversalStack(bs))
}

func main() {
	run()
}
