package bstree

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Print(node *TreeNode) {
	if node == nil {
		return
	}
        fmt.Print(node.Val)
	fmt.Print(" left: ")
	if node.Left != nil {
		fmt.Print(node.Left.Val)
	} else {
		fmt.Print(node.Left)
	}
	fmt.Print(" right: ")
	if node.Right != nil {
		fmt.Println(node.Right.Val)
	} else {
		fmt.Println(node.Right)
	}
	Print(node.Left)
	Print(node.Right)	
}

func NewTree(array []int) *TreeNode {
	var root *TreeNode
	for _, val := range array {
		root = addNode(root, val)
	}
	return root
}

func addNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if node.Val > val {
		node.Left = addNode(node.Left, val)
	} else {
		node.Right = addNode(node.Right, val)
	}
	return node
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%d\n", t.Val)
}