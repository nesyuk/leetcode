type TreeNode struct {
    Val int
    Size int
    Left *TreeNode
    Right *TreeNode
}

func (n *TreeNode) Add(val int) *TreeNode {
    if n == nil {
        return &TreeNode{Val: val, Size: 1}
    }
    if n.Val <= val {
        n.Right = n.Right.Add(val)
    } else {
        n.Left = n.Left.Add(val)
    }
    n.Size++
    return n
}

func (n *TreeNode) findKth(k int) int {
    greaterThanNode := 0
    if n.Right != nil {
        greaterThanNode = n.Right.Size
    }
    if k == greaterThanNode + 1 {
        return n.Val
    }
    if k <= greaterThanNode {
        return n.Right.findKth(k)
    }
    if n.Left != nil {
        return n.Left.findKth(k - greaterThanNode - 1)
    }
    return -1
}

type KthLargest struct {
    root *TreeNode
    k int
}


func Constructor(k int, nums []int) KthLargest {
    kthLargest := &KthLargest{k: k}
    for _, num := range nums {
        kthLargest.Add(num)
    }
    return *kthLargest
}


func (this *KthLargest) Add(val int) int {
    this.add(val)
    return this.findKthLargest()
}

func (this *KthLargest) add(val int) {
    if this.root == nil {
        this.root = &TreeNode{Val: val}
        return
    }
    this.root.Add(val)
}

func (this *KthLargest) findKthLargest() int {
    return this.root.findKth(this.k)
}
