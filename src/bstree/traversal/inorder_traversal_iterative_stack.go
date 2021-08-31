package main

import (
	"fmt"
	"bstree"
)

func inorderTraversal(root *bstree.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*bstree.TreeNode{}
	current := root
	for len(stack) != 0 || current != nil {
		if current != nil {
			// go down
			stack = append(stack, current)
			current = current.Left
		} else {
			// go up
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			current = current.Right			
		}
	}
	return result
}

func run() {
	fmt.Println("go!")
	bs := bstree.NewTree([]int{2, 1, 3})
	fmt.Println(inorderTraversal(bs))

	bs = bstree.NewTree([]int{5, 6, 3, 1, 5, 7, 8, 9, 2})
	fmt.Println(inorderTraversal(bs))

}

func main() {
	run()
}
