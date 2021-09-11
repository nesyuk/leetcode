package bstree

func InorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    if p.Right != nil {
        return leftMostChild(p.Right)
    }
    return rightMostParent(root, p)
}

func leftMostChild(node *TreeNode) *TreeNode {
    child := node
    for child.Left != nil {
        child = child.Left
    }
    return child
}

func rightMostParent(node *TreeNode, p *TreeNode) *TreeNode {
    var parent *TreeNode
    current := node
    for current != nil && current != p {
        if current.Val > p.Val {
            parent = current
            current = current.Left
        } else {
            current = current.Right
        }
    }
    return parent
}
