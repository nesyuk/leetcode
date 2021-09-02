package main

import (
	"fmt"
	"bstree"
)

func inorderTraversalMorris(root *bstree.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	node := root
	for node != nil {
		if node.Left == nil {
			result = append(result, node.Val)
			node = node.Right
		} else {
			rightMost := node.Left
			for rightMost.Right != nil && rightMost.Right != node {
				rightMost = rightMost.Right
			}
			if rightMost.Right == nil {
				rightMost.Right = node
				node = node.Left
			} else {
				// restore the tree
				rightMost.Right = nil
				result = append(result, node.Val)
				node = node.Right
			}
		}
	}
	return result
}

func run() {
	fmt.Println("go!")
	bs := bstree.NewBSTree([]int{2, 1, 3})
	fmt.Println(inorderTraversalMorris(bs))

	bs = bstree.NewBSTree([]int{5, 6, 3, 1, 5, 7, 8, 9, 2})
	fmt.Println(inorderTraversalMorris(bs))

	bs = bstree.NewBSTree([]int{1, 3, 2})
	fmt.Println(inorderTraversalMorris(bs))
}

func main() {
	run()
}
