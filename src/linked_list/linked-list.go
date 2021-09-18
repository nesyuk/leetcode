package linked_list

type ListNode struct {
    Val int
    Next *ListNode
}

func NewLinkedListWithCycle(vals []int, pos int) *ListNode {
        root := &ListNode{}
        node := root
        for _, val := range vals {
                node.Next = &ListNode{Val: val}
                node = node.Next
        }
	current := root.Next
	fromNode := current
	toNode := current
	for count := 0; current != nil; count++ {
		if count == pos {
			fromNode = current
		}
		toNode = current
		current = current.Next
	}
	toNode.Next = fromNode
        return root.Next
}

func NewLinkedList(vals []int) *ListNode {
	root := &ListNode{}
	node := root
	for _, val := range vals {
		node.Next = &ListNode{Val: val}
		node = node.Next
	}
	return root.Next
}
type MyLinkedList struct {
    head *ListNode
    size int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || this.size <= index {
        return -1
    }
    current := this.head
    for count := 0; count != index; count++ {
        current = current.Next
    }
    return current.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    this.head = &ListNode{Val: val, Next: this.head}
    this.size++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    node := &ListNode{Val: val}
    
    if this.head == nil {
        this.head = node
    } else {
        tail := this.head
        for tail.Next != nil {
            tail = tail.Next
        }
        tail.Next = node
    }
    this.size++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.size {
        return
    }
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    current := this.head
    for count := 0; count < index - 1; count++ {
        current = current.Next
    }
    next := current.Next
    current.Next = &ListNode{Val: val, Next: next}
    this.size++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }
    if index == 0 {
        this.head = this.head.Next
    } else {
        current := this.head
        for count := 0; count < index - 1; count++ {
            current = current.Next
        }
        current.Next = current.Next.Next
    }
    this.size--
}
