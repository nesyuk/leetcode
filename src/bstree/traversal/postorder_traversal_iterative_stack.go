package main

import (
	"fmt"
	"bstree"
)

func postorderTraversalStack(root *bstree.TreeNode) []int {
	result := []int{}
	stack := []*bstree.TreeNode{}
	current := root
	for {
		for current != nil {
			stack = append(stack, current)
			stack = append(stack, current)
			current = current.Left
		}

		if len(stack) == 0 {
			return result
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) != 0 && stack[len(stack)-1] == current {
			current = current.Right
		} else {
			result = append(result, current.Val)
			current = nil
		}
	}
	return result
}

func run() {
	fmt.Println("go!")
	bs := bstree.NewTree([]int{1, 2, 3, 5})
	fmt.Println(postorderTraversalStack(bs))

	bs = bstree.NewTree([]int{5, 6, 3, 1, 5, 7, 8, 9, 2})
	fmt.Println(postorderTraversalStack(bs))
}

func main() {
	run()
}
