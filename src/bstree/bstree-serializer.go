package bstree

import (
    "fmt"
    "strconv"
    "strings"
)

type Codec struct {
    separator string
}

func NewCodec() Codec {
    return Codec{separator: ","}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    stack := []*TreeNode{root}
    encodedTree := ""
    levelLen := 1
    for len(stack) != 0 {
        levelLen = len(stack)
        for i := 0; i < levelLen; i++ {
            current := stack[0]
            stack = stack[1:]
            if current == nil {
                encodedTree = fmt.Sprintf("%s%s", encodedTree, c.separator)
            } else {
                encodedTree = fmt.Sprintf("%s%s%d", encodedTree, c.separator, current.Val)
                stack = append(stack, current.Left)
                stack = append(stack, current.Right)
            }
        }
    }
    // remove trailing nulls
    encodedTree = encodedTree[:len(encodedTree) - levelLen]
    return strings.Replace(encodedTree, c.separator, "", 1)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {  
    if len(data) == 0 {
        return nil
    }
    tokens := strings.Split(data, c.separator)
    root := c.decode(tokens[0])
    stack := []*TreeNode{root}
    idx := 0
    for len(stack) != 0 {
        current := stack[0]
        stack = stack[1:]
        idx++
        if idx < len(tokens) {
            current.Left = c.decode(tokens[idx])
            if current.Left != nil {
                stack = append(stack, current.Left)
            }
        }
        idx++
        if idx < len(tokens) {
            current.Right = c.decode(tokens[idx])
            if current.Right != nil {
                stack = append(stack, current.Right)
            }
        }
    }
    return root
}

func (c *Codec) decode(str string) *TreeNode {
    if len(str) == 0 {
        return nil
    }
    val, _ := strconv.Atoi(str)
    return &TreeNode{Val: val}
}
